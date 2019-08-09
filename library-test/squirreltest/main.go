package main

import (
	"log"

	"roshar/services/rewards-svc/pkg/datastore/pg"
	"roshar/services/rewards-svc/pkg/rewards"
	"roshar/services/users-svc/pkg/users"
	"roshar/utils/structutil"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type NearbyCampaign struct {
	Lat       float64              `db:"lat" protobuf:"name=lat"`
	Lng       float64              `db:"lng" protobuf:"name=lng"`
	Quantity  int32                `db:"quantity" protobuf:"name=quantity"`
	Title     string               `db:"title" protobuf:"name=title"`
	Counter   uint64               `db:"counter" protobuf:"name=counter"`
	ExpiredAt time.Time            `db:"expired_at" protobuf:"name=expired_at"`
	Merchant  *rewards.Merchant    `db:"merchants" protobuf:"name=merchant"`
	Reward    *rewards.RewardStamp `db:"reward_stamps" protobuf:"name=reward"`
}

func main() {
	saveBlocks()
}

func saveBlocks() {
	ds, err := pg.NewDatastore("postgres://:@localhost:5432/gc_wallets_development?sslmode=disable")
	if err != nil {
		panic(err)
	}

	now := ds.Timestamp()

	stmt, args, err := ds.SQLBuilder.
		Insert("ethereum_blocks").
		Columns("number", "created_at", "updated_at").
		Values(5277349, now, now).
		Values(5277347, now, now).
		Suffix("ON CONFLICT DO NOTHING").
		ToSql()

	if err != nil {
		log.Println(err)
	}

	log.Println(stmt)

	_, err = ds.Exec(stmt, args...)
	if err != nil {
		log.Println(err)
	}
}

func missingBlocks() {
	ds, err := pg.NewDatastore("postgres://:@localhost:5432/gc_wallets_development?sslmode=disable")
	if err != nil {
		panic(err)
	}

	stmt := `
		SELECT s.i
		FROM generate_series($1::int, $2::int) s(i)
		WHERE NOT EXISTS (
			SELECT 1
			FROM ethereum_blocks
			WHERE number = s.i
		)
	`

	var blocks []int64
	start := int64(4880000)
	end := int64(5324091)
	err = ds.Select(&blocks, stmt, start, end)
	if err != nil {
		log.Println(err)
	}

	log.Println(blocks)
}

func userStats() {
	db, err := pg.NewDatastore("postgres://:@localhost:5432/gc_users_development?sslmode=disable")
	if err != nil {
		panic(err)
	}

	stmt, args, err := db.SQLBuilder.
		Select(
			"count(*) total",
			"sum(case when users.updated_at between current_date::timestamp and current_date::timestamp + interval '1 day' - interval '1 second' then 1 else 0 end) \"new\"",
			"sum(case when gender = 'male' then 1 else 0 end) male",
			"sum(case when gender = 'female' then 1 else 0 end) female",
			"coalesce(trunc(avg(case when gender = 'male' then extract(years from age(now(), dob)) end)), 0) male_avg_age",
			"coalesce(trunc(avg(case when gender = 'female' then extract(years from age(now(), dob)) end)), 0) female_avg_age",
		).
		From("users").
		LeftJoin("profiles ON users.id = profiles.user_id").
		Where(sq.Eq{"status": "active"}).
		ToSql()

	if err != nil {
		log.Printf("sql error: %s", err)
		return
	}

	// var count int
	// err = db.QueryRowx(stmt, args...).Scan(&count)
	// if err != nil {
	// 	log.Printf("get error: %s", err)
	// 	return
	// }
	// log.Println("count", count)

	stats := &users.UserStats{}
	err = db.Get(stats, stmt, args...)
	if err != nil {
		log.Printf("get error: %s", err)
		return
	}
	log.Println("stats", stats)
}

func selectListings() {
	ds, err := pg.NewDatastore("postgres://:@localhost:5432/gc_rewards_development?sslmode=disable")
	if err != nil {
		panic(err)
	}

	listing := &rewards.Listing{}
	listingFields := structutil.GenerateDBSelectFields(listing)

	stmt, args, err := ds.SQLBuilder.
		Select(listingFields,
			"reward_cards.id \"reward_cards.id\"", "reward_cards.title \"reward_cards.title\"",
			"reward_cards.status \"reward_cards.status\"", "starred \"reward_cards.starred\"",
			"reward_stamps.images \"reward_cards.images\"", "reward_stamps.expired_at \"reward_cards.expired_at\"",
			"merchants.id \"reward_cards.merchant_id\"", "merchants.name \"reward_cards.name\"").
		From("listings").
		Join("reward_cards ON listings.reward_card_id = reward_cards.id").
		Join("merchants ON reward_cards.merchant_id = merchants.id").
		Join("reward_stamps ON reward_cards.reward_stamp_id = reward_stamps.id").
		Where(
			sq.GtOrEq{"expired_at": ds.Timestamp()},
		).
		ToSql()

	if err != nil {
		log.Println(err)
	}

	listings := make([]rewards.ListingSummary, 0)
	err = ds.Select(&listings, stmt, args...)
	if err != nil {
		log.Println("select", err)
	}

	log.Println(listings[0].Listing)
	log.Println(listings[0].RewardCard)
	log.Println(listings)

}

func updateInterest() {
	// db, err := pg.NewDatastore("postgres://:@localhost:5432/gc_users_development?sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }

	// stmt, args, err := db.SQLBuilder.
	// 	Select("interests").
	// 	From("profiles").
	// 	Where(sq.Eq{"user_id": 14}).
	// 	ToSql()

	// u := users.Profile{}

	// interests := []string{}

	// err = db.Get(pq.Array(&interests), stmt, args...)
	// log.Println(err)
	// log.Println(interests)
	// log.Println(u)
}

func update() {
	db, err := pg.NewDatastore("postgres://:@localhost:5432/gc_rewards_development?sslmode=disable")
	if err != nil {
		panic(err)
	}

	a := &rewards.RewardCardDetail{}

	stmt, args, err := db.SQLBuilder.
		Update("reward_cards").
		Set("updated_at", time.Now().UTC()).
		Where(sq.Eq{"id": 100}).
		Suffix("RETURNING *").
		ToSql()

	log.Println(err)

	// _, err = db.Exec(stmt, args...)
	// err = db.QueryRow(stmt, args...).Scan(a)
	// err = db.Get(&a, stmt, args...)
	err = db.QueryRowx(stmt, args...).Scan(a)

	log.Println(err)
	log.Println(a)
}

func selectCampaigns(db *pg.Datastore) {

	gender := "all"
	lat := 39.142293
	lng := -105.583300
	radius := 97999.01
	age := -1
	expiredAt := time.Now().UTC()
	userID := 14

	merchant := &rewards.Merchant{}
	merchantFields := structutil.GenerateDBSelectFields(merchant)

	stamp := &rewards.RewardStamp{}
	stampFields := structutil.GenerateDBSelectFields(stamp)

	stmt, args, err := db.SQLBuilder.
		Select("lat", "lng", "quantity", "counter", "c.title AS title",
			merchantFields, stampFields).
		From("campaigns AS c").
		Join("merchants ON c.merchant_id = merchants.id").
		Join("campaign_locations AS cl ON c.id = cl.campaign_id").
		Join("reward_stamps ON c.id = reward_stamps.campaign_id").
		Where("earth_distance(ll_to_earth(cl.lat, cl.lng), ll_to_earth(?, ?)) < sec_to_gc(cl.radius) + sec_to_gc(?)", lat, lng, radius).
		Where("c.counter < c.quantity").
		Where(
			sq.Case().
				When(sq.Eq{"c.gender": "all"}, "TRUE").
				Else(sq.Eq{"c.gender": gender})).
		Where(
			sq.Or{
				sq.Expr("? = ?", age, -1),
				sq.And{
					sq.Case("c.min_age").When("NULL", "TRUE").Else(sq.LtOrEq{"c.min_age": age}),
					sq.Case("c.max_age").When("NULL", "TRUE").Else(sq.GtOrEq{"c.max_age": age}),
				},
			}).
		Where(sq.And{
			sq.Eq{"c.status": "active"},
			sq.GtOrEq{"reward_stamps.expired_at": expiredAt},
		}).
		Where("NOT EXISTS (SELECT * FROM reward_cards as r WHERE r.campaign_id = c.id AND r.origin_owner_id = ?)", userID).
		ToSql()

	campaigns := []NearbyCampaign{}
	err = db.Select(&campaigns, stmt, args...)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(stmt)
	log.Println(args)
	log.Println(err)

	log.Println(campaigns)
	log.Println(campaigns[0].Merchant)
	log.Println(campaigns[0].Reward)
}
