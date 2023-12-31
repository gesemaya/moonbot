// Code generated by squick at 2021-07-21T14:56:47+03:00.
// squick make -table votes insert get:done set:updated_at,message_id,days_left,done
package database

import (
	"reflect"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
)

type Vote struct {
	CreatedAt time.Time     `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time     `db:"updated_at" json:"updatedAt"`
	ID        int           `db:"id" json:"id"`
	Done      bool          `db:"done" json:"done"`
	DaysLeft  int           `db:"days_left" json:"daysLeft"`
	Ideas     pq.Int32Array `db:"ideas" json:"ideas"`
	MessageID string        `db:"message_id" json:"messageID"`

	db *DB
}

func (db *DB) InsertVote(vote Vote) (id int, _ error) {
	data := map[string]interface{}{
		"created_at": vote.CreatedAt,
		"updated_at": vote.UpdatedAt,
		"id":         vote.ID,
		"done":       vote.Done,
		"days_left":  vote.DaysLeft,
		"ideas":      vote.Ideas,
		"message_id": vote.MessageID,
	}
	for _, col := range voteColumns {
		if reflect.ValueOf(data[col]).IsZero() {
			delete(data, col)
		}
	}

	var (
		cols []string
		vals []interface{}
	)
	for col, val := range data {
		cols = append(cols, col)
		vals = append(vals, val)
	}

	query, args, err := squirrel.
		Insert("votes").
		Columns(cols...).
		Values(vals...).
		Suffix(`RETURNING "id"`).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return id, err
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		return id, err
	}
	if rows.Next() {
		rows.Scan(&id)
	}

	return
}

func (db *DB) VoteByDone(done bool) (vote Vote, _ error) {
	vote.db = db
	const query = `SELECT * FROM votes WHERE done=$1`
	return vote, db.Get(&vote, query, done)
}

func (vote *Vote) SetUpdatedAt(updatedAt time.Time) error {
	vote.UpdatedAt = updatedAt
	const query = `UPDATE votes SET updated_at=$1 WHERE id=$2`
	_, err := vote.db.Exec(query, updatedAt, vote.ID)
	return err
}

func (vote *Vote) SetMessageID(messageID string) error {
	vote.MessageID = messageID
	const query = `UPDATE votes SET message_id=$1 WHERE id=$2`
	_, err := vote.db.Exec(query, messageID, vote.ID)
	return err
}

func (vote *Vote) SetDaysLeft(daysLeft int) error {
	vote.DaysLeft = daysLeft
	const query = `UPDATE votes SET days_left=$1 WHERE id=$2`
	_, err := vote.db.Exec(query, daysLeft, vote.ID)
	return err
}

func (vote *Vote) SetDone(done bool) error {
	vote.Done = done
	const query = `UPDATE votes SET done=$1 WHERE id=$2`
	_, err := vote.db.Exec(query, done, vote.ID)
	return err
}

var voteColumns = []string{
	"created_at",
	"updated_at",
	"id",
	"done",
	"days_left",
	"ideas",
	"message_id",
}
