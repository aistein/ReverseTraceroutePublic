/*
Copyright (c) 2015, Northeastern University
 All rights reserved.

 Redistribution and use in source and binary forms, with or without
 modification, are permitted provided that the following conditions are met:
     * Redistributions of source code must retain the above copyright
       notice, this list of conditions and the following disclaimer.
     * Redistributions in binary form must reproduce the above copyright
       notice, this list of conditions and the following disclaimer in the
       documentation and/or other materials provided with the distribution.
     * Neither the name of the Northeastern University nor the
       names of its contributors may be used to endorse or promote products
       derived from this software without specific prior written permission.

 THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 DISCLAIMED. IN NO EVENT SHALL Northeastern University BE LIABLE FOR ANY
 DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package sql

import (
	"database/sql"
	"fmt"
	"time"

	dm "github.com/NEU-SNS/ReverseTraceroute/datamodel"
)
import "github.com/go-sql-driver/mysql"

type DB struct {
	db *sql.DB
}

type DbConfig struct {
	UName    string
	Password string
	Host     string
	Port     string
	Db       string
}

var conFmt string = "%s:%s@tcp(%s:%s)/%s?parseTime=true"

func NewDB(con DbConfig) (*DB, error) {
	conString := fmt.Sprintf(conFmt, con.UName, con.Password, con.Host, con.Port, con.Db)
	db, err := sql.Open("mysql", conString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(24)
	return &DB{db: db}, nil
}

type VantagePoint struct {
	Ip           uint32
	Controller   sql.NullInt64
	HostName     string
	Site         string
	TimeStamp    bool
	RecordRoute  bool
	CanSpoof     bool
	Active       bool
	ReceiveSpoof bool
	LastUpdated  time.Time
	SpoofChecked mysql.NullTime
}

func (vp *VantagePoint) ToDataModel() *dm.VantagePoint {
	nvp := &dm.VantagePoint{}
	nvp.Ip = vp.Ip
	if vp.Controller.Valid {
		nvp.Controller = uint32(vp.Controller.Int64)
	}
	nvp.Hostname = vp.HostName
	nvp.Timestamp = vp.TimeStamp
	nvp.RecordRoute = vp.TimeStamp
	nvp.CanSpoof = vp.CanSpoof
	nvp.Active = vp.Active
	nvp.ReceiveSpoof = vp.ReceiveSpoof
	nvp.LastUpdated = vp.LastUpdated.Unix()
	nvp.Site = vp.Site
	if vp.SpoofChecked.Valid {
		nvp.SpoofChecked = vp.LastUpdated.Unix()
	}
	return nvp
}

const (
	getVpsQuery string = `
SELECT
	ip, controller, hostname, timestamp,
	record_route, can_spoof, active,
    receive_spoof, last_updated
FROM
	vantage_point;
`
)

func (db *DB) GetVPs() ([]*dm.VantagePoint, error) {
	rows, err := db.db.Query(getVpsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	vps := make([]*dm.VantagePoint, 0)
	for rows.Next() {
		vp := &VantagePoint{}
		err := rows.Scan(
			&vp.Ip,
			&vp.Controller,
			&vp.HostName,
			&vp.TimeStamp,
			&vp.RecordRoute,
			&vp.CanSpoof,
			&vp.Active,
			&vp.ReceiveSpoof,
			&vp.LastUpdated,
		)
		if err != nil {
			return vps, err
		}
		vps = append(vps, vp.ToDataModel())
	}
	err = rows.Err()
	return vps, err
}

const (
	updateControllerQuery string = `
UPDATE
	vantage_point
SET
	controller = ?,
	active = ?
WHERE
	ip = ?
`
	updateControllerQueryNull string = `
UPDATE
	vantage_point
SET
	controller = IF(controller = ?, NULL, controller),
	active = IF(controller = ?, ?, active)
WHERE
	ip = ?
`
)

func (db *DB) UpdateController(ip, newc, con uint32) error {
	args := make([]interface{}, 0)
	var active bool
	query := updateControllerQuery
	if newc == 0 {
		query = updateControllerQueryNull
		args = append(args, con, con)
	} else {
		args = append(args, newc)
		active = true
	}
	args = append(args, active, ip)
	_, err := db.db.Exec(query, args...)
	return err
}

const (
	updateActiveQuery string = `
UPDATE
	vantage_point
SET
	active = ?
WHERE
	ip = ?
`
)

func (db *DB) UpdateActive(ip uint32, active bool) error {
	_, err := db.db.Exec(updateActiveQuery, active, ip)
	return err
}

const (
	updateCanSpoofQuery string = `
UPDATE
	vantage_point
SET
	can_spoof = ?,
	spoof_checked = NOW()
WHERE
	ip = ?
`
)

func (db *DB) UpdateCanSpoof(ip uint32, canSpoof bool) error {
	_, err := db.db.Exec(updateCanSpoofQuery, canSpoof, ip)
	return err
}

func (db *DB) Close() error {
	return db.db.Close()
}