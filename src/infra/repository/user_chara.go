package repository

import (
	"context"
	"database/sql"

	"go-gacha-server/src/domain/entity"
	"go-gacha-server/src/domain/repository"
)

type userCharaRepository struct {
	db *sql.DB
}

func NewUserCharaRepository(db *sql.DB) repository.UserCharaRepository {
	return &userCharaRepository{db: db}
}

func (ucr *userCharaRepository) List(ctx context.Context, ue entity.User) ([]*entity.UserChara, error) {
	const list = `SELECT user_character_possessions.id, characters.id, characters.name, characters.icon_url, rarities.rarity, rarities.probability FROM user_character_possessions INNER JOIN characters ON user_character_possessions.character_id = characters.id INNER JOIN rarities ON characters.rarity_id = rarities.id WHERE user_character_possessions.user_id = ?`

	stmt, err := ucr.db.PrepareContext(ctx, list)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, ue.Id)
	if err != nil {
		return nil, err
	}

	var entities []*entity.UserChara
	for rows.Next() {
		uce := &entity.UserChara{User: ue}
		ce := &entity.Chara{}

		err := rows.Scan(&uce.Id, &ce.Id, &ce.Name, &ce.IconURL, &ce.Rarity, &ce.Probability)
		if err != nil {
			return nil, err
		}

		uce.Chara = *ce
		entities = append(entities, uce)
	}

	return entities, nil
}

func (ucr *userCharaRepository) Store(ctx context.Context, ue entity.User, ces []*entity.Chara) error {
	const store = `INSERT user_character_possessions (user_id, character_id) VALUE (?, ?)`

	stmt, err := ucr.db.PrepareContext(ctx, store)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, ce := range ces {
		_, err := stmt.ExecContext(ctx, ue.Id, ce.Id)
		if err != nil {
			return err
		}
	}

	return nil
}
