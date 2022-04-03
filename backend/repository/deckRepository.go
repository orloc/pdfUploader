package repository

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"wefunder/backend/entity"
)

type DeckRepository struct {
	db *pgxpool.Pool
	ctx context.Context
	tName string
}

func NewDeckRepository(ctx context.Context, db *pgxpool.Pool) *DeckRepository {
	d := new(entity.Deck)
	return &DeckRepository{
		db:  db,
		ctx: ctx,
		tName: d.TableName(),
	}
}

func (r *DeckRepository) Exists(name string)  (*entity.Deck, error) {
	q := getQueryBuilder().
		Select("*").
		From(r.tName).
		Where(sq.Eq{
			"company_name": name,
		})

	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := execQuery(r.ctx, r.db, sql, args, r.getRows)
	if rows != nil {
		dat := rows.([]*entity.Deck)
		var res *entity.Deck
		if len(dat) > 0 {
			res = dat[0]
		}
		return res, nil
	}

	return nil, nil
}

func (r *DeckRepository) LoadDecks()  ([]*entity.Deck, error) {
	q := getQueryBuilder().
		Select("*").
		From(r.tName).
		OrderBy("created_at")
	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := execQuery(r.ctx, r.db, sql, args, r.getRows)
	if rows != nil  {
		return rows.([]*entity.Deck), nil
	}

	return []*entity.Deck{}, nil
}

func (r *DeckRepository) CreateDeck(dat *entity.Deck)  error {
	sql, args, err := getQueryBuilder().
		Insert(r.tName).
		Columns("company_name", "images", "uuid").
		Values(dat.CompanyName, dat.Images, dat.Uuid).ToSql()

	if err != nil {
		return err
	}
	_, err = execQuery(r.ctx, r.db, sql, args, r.getRows)
	return nil
}

func (r *DeckRepository) getRows(queryRows pgx.Rows) (interface{}, error) {
	var results []*entity.Deck
	if scanErr := pgxscan.ScanAll(&results, queryRows); scanErr != nil {
		return nil, scanErr
	}
	return results, nil
}
