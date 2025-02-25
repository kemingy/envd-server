// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: image.sql

package query

import (
	"context"

	"github.com/jackc/pgtype"
)

const createImageInfo = `-- name: CreateImageInfo :one
INSERT INTO image_info (
  owner_token, name, digest, created, size, labels
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id, owner_token, name, digest, created, size, labels
`

type CreateImageInfoParams struct {
	OwnerToken string       `json:"owner_token"`
	Name       string       `json:"name"`
	Digest     string       `json:"digest"`
	Created    int64        `json:"created"`
	Size       int64        `json:"size"`
	Labels     pgtype.JSONB `json:"labels"`
}

func (q *Queries) CreateImageInfo(ctx context.Context, arg CreateImageInfoParams) (ImageInfo, error) {
	row := q.db.QueryRow(ctx, createImageInfo,
		arg.OwnerToken,
		arg.Name,
		arg.Digest,
		arg.Created,
		arg.Size,
		arg.Labels,
	)
	var i ImageInfo
	err := row.Scan(
		&i.ID,
		&i.OwnerToken,
		&i.Name,
		&i.Digest,
		&i.Created,
		&i.Size,
		&i.Labels,
	)
	return i, err
}

const getImageInfo = `-- name: GetImageInfo :one
SELECT id, owner_token, name, digest, created, size, labels FROM image_info
WHERE owner_token = $1 AND name = $2 LIMIT 1
`

type GetImageInfoParams struct {
	OwnerToken string `json:"owner_token"`
	Name       string `json:"name"`
}

func (q *Queries) GetImageInfo(ctx context.Context, arg GetImageInfoParams) (ImageInfo, error) {
	row := q.db.QueryRow(ctx, getImageInfo, arg.OwnerToken, arg.Name)
	var i ImageInfo
	err := row.Scan(
		&i.ID,
		&i.OwnerToken,
		&i.Name,
		&i.Digest,
		&i.Created,
		&i.Size,
		&i.Labels,
	)
	return i, err
}

const listImageByOwner = `-- name: ListImageByOwner :many
SELECT id, owner_token, name, digest, created, size, labels FROM image_info
WHERE owner_token = $1
`

func (q *Queries) ListImageByOwner(ctx context.Context, ownerToken string) ([]ImageInfo, error) {
	rows, err := q.db.Query(ctx, listImageByOwner, ownerToken)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ImageInfo
	for rows.Next() {
		var i ImageInfo
		if err := rows.Scan(
			&i.ID,
			&i.OwnerToken,
			&i.Name,
			&i.Digest,
			&i.Created,
			&i.Size,
			&i.Labels,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
