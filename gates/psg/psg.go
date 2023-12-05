package psg

import (
	"HW1_http/pkg"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/url"
)

type Psg struct {
	conn *pgxpool.Pool
}

func NewPsg(dburl string, login, pass string) (psg *Psg, err error) {
	eW := pkg.NewEWrapper("NewPsg()")

	psg = &Psg{}
	psg.conn, err = parseConnectionString(dburl, login, pass)
	if err != nil {
		err = eW.WrapError(err, "parseConnectionString(dburl, login, pass)")
		return nil, err
	}

	err = psg.conn.Ping(context.Background())
	if err != nil {
		err = eW.WrapError(err, "psg.conn.Ping(context.Background())")
		return nil, err
	}

	return
}

func parseConnectionString(dburl, user, password string) (db *pgxpool.Pool, err error) {
	eW := pkg.NewEWrapper("parseConnectionString()")

	var u *url.URL
	if u, err = url.Parse(dburl); err != nil {
		err = eW.WrapError(err, "url.Parse(dburl)")
		return nil, err
	}
	u.User = url.UserPassword(user, password)
	db, err = pgxpool.New(context.Background(), u.String())
	if err != nil {
		err = eW.WrapError(err, "pgxpool.New(context.Background(), u.String())")
		return nil, err
	}
	return
}
