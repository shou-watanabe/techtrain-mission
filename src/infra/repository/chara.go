package repository

import (
	"context"
	"database/sql"

	"go-gacha-server/src/domain/entity"
	"go-gacha-server/src/domain/repository"
)

type charaRepository struct {
	db *sql.DB
}

func NewCharaRepository(db *sql.DB) repository.CharaRepository {
	return &charaRepository{db: db}
}

func (cr *charaRepository) List(ctx context.Context) ([]*entity.Chara, error) {
	const list = `SELECT characters.id, characters.name, characters.icon_url, rarities.rarity, rarities.probability FROM characters, rarities where characters.rarity_id = rarities.id`

	stmt, err := cr.db.PrepareContext(ctx, list)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}

	var entities []*entity.Chara
	for rows.Next() {
		ce := &entity.Chara{}

		err := rows.Scan(&ce.Id, &ce.Name, &ce.IconURL, &ce.Rarity, &ce.Probability)
		if err != nil {
			return nil, err
		}

		entities = append(entities, ce)
	}

	return entities, nil
}
