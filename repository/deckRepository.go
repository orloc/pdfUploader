package repository

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"wefunder/entity"
)

type DeckRepository struct {
	db *pgxpool.Pool
	ctx context.Context
}

func NewDeckRepository(ctx context.Context, db *pgxpool.Pool) *DeckRepository {
	return &DeckRepository{
		db:  db,
		ctx: ctx,
	}
}

func (r *DeckRepository) LoadDecks()  ([]*entity.Deck, error) {
	d := new(entity.Deck)
	q := getQueryBuilder().
		Select("*").
		From(d.TableName()).
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

func (r *DeckRepository) CreateMessage(dat *entity.Deck)  error {
	sql, args, err := getQueryBuilder().
		Insert(dat.TableName()).
		Columns("company_name", "images").
		Values(dat.CompanyName, dat.Images).ToSql()

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
