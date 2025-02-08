package drivers

import (
	"database/sql"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"
	_ "modernc.org/sqlite"
)

type SqlLiteDatabase struct {
	cursor       *sql.DB
	DatabasePath string
}

func (db *SqlLiteDatabase) InitDatabase(force bool) error {
	if db.cursor != nil && !force {
		log.Info("Database Already Inizialized")
		return nil
	}

	if db.DatabasePath == "" {
		log.Error("Please make sure to define a path")
		return errors.New("no path defined")
	}
	database, err := sql.Open("sqlite", db.DatabasePath)

	if err != nil {
		log.Error(err)
		return errors.New("couldn't inizialize db")
	}

	db.cursor = database
	transaction, err := db.cursor.Begin()
	if err != nil {
		log.Fatal(err)
		return errors.New("Cannot begin transaction")
	}
	_, err = transaction.Exec(CreateManwhaResourceTable)

	if err != nil {
		log.Warning(err)
	}

	_, err = transaction.Exec(CreateManwhaPageTable)

	if err != nil {
		log.Warning(err)
	}

	_, err = transaction.Exec(CreateManwhaPageURLTable)

	if err != nil {
		log.Warning(err)
	}

	err = transaction.Commit()

	if err != nil {
		log.Fatal(err)
		return errors.New("Cannot commit transaction")
	}

	return nil
}

func (db *SqlLiteDatabase) AddManwhaResource(resource *ManwhaResource, driver string, addPages bool) error {
	err := db.InitDatabase(false)
	if err != nil {
		log.Fatal("Cannot Add AddManwhaResource database cannot be inizialized")
		return err
	}

	res, err := db.cursor.Exec(InsertManwhaResource, resource.name, resource.address, resource.nChapter, resource.imageUrl, driver)

	if err != nil {
		log.Fatal(err)
		return errors.New("cannot Insert ManwhaPage into db")
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Cannot get last inserted Id of mawna resource ")
		return err
	}

	resource.id = int(id)
	if !addPages {
		return nil
	}

	for i, page := range resource.pages {
		err = db.AddManwhaPage(page, resource.id)
		if err != nil {
			log.Warnf("Cannot add to database %d of %s", i, resource.name)
		}
	}
	return nil
}

func (db *SqlLiteDatabase) AddManwhaPage(page *ManwhaPage, resourceId int) error {
	err := db.InitDatabase(false)
	if err != nil {
		log.Fatal("Cannot Add AddManwhaResource database cannot be inizialized")
		return err
	}
	transaction, err := db.cursor.Begin()

	if err != nil {
		log.Fatal(err)
		return errors.New("cannot start the transaction")
	}

	res, err := transaction.Exec(InsertManwhaPage, page.pageNumber, resourceId)

	if err != nil {
		log.Fatal(err)
		return errors.New("cannot Insert ManwhaPage into db")
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Cannot get last inserted Id of mawna resource ")
		return err
	}

	page.id = int(id)

	for i, url := range page.ImageUrls {
		_, err := transaction.Exec(InsertManwhaPageURL, url, page.id)
		if err != nil {
			log.Fatal(err)
			return fmt.Errorf("cannot insert url %d into the database", i)
		}

	}

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
		return errors.New("Cannot commit transaction on page number")
	}
	return nil

}

func (db *SqlLiteDatabase) GetManwhaResource(id int, collectPages bool) (*ManwhaResource, error) {
	err := db.InitDatabase(false)
	if err != nil {
		log.Fatal("Cannot GET ManwhaResource,database cannot be inizialized")
		return nil, err
	}

	row, err := db.cursor.Query(SelectManwhaPageById, id)

	if err != nil {
		log.Fatal("Error while querying")
		return nil, err
	}

	hasRow := row.Next()

	if !hasRow {
		log.Info("No Manwha resource with %d found", id)

		return nil, nil
	}

	resource := new(ManwhaResource)
	err = row.Scan(resource)

	if err != nil {
		log.Fatal(err)
		return nil, errors.New("cannot scan row inside resource")
	}

	if !collectPages {
		return resource, nil
	}

	page, err := db.GetManwhaPageByResourceId(resource.id)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("cannot retrieve page by resourceId")
	}
	resource.pages = append(resource.pages, page)
	return resource, nil
}

func (db *SqlLiteDatabase) GetManwhaPage(id int) (*ManwhaPage, error) {
	err := db.InitDatabase(false)
	if err != nil {
		log.Fatal("Cannot GET ManwhaResource,database cannot be inizialized")
		return nil, err
	}

	rows, err := db.cursor.Query(JoinManwhaPage, id)
	if err != nil {
		log.Fatal("Error while querying")
		return nil, err
	}
	manwhaPage := new(ManwhaPage)
	urls := make([]string, 0)
	for rows.Next() {
		var currStr string
		err = rows.Scan(&manwhaPage.id, &manwhaPage.pageNumber, &currStr)
		if err != nil {
			log.Warn("cannot scan url %s", err.Error())
			continue
		}
		urls = append(urls, currStr)
	}
	manwhaPage.ImageUrls = urls

	return manwhaPage, nil
}

func (db *SqlLiteDatabase) GetManwhaPageByResourceId(resourceId int) (*ManwhaPage, error) {
	err := db.InitDatabase(false)
	if err != nil {
		log.Fatal("Cannot GET ManwhaResource,database cannot be inizialized")
		return nil, err
	}

	rows, err := db.cursor.Query(JoinManwhaPageByResourceId, resourceId)
	if err != nil {
		log.Fatal("Error while querying")
		return nil, err
	}
	manwhaPage := new(ManwhaPage)
	urls := make([]string, 0)
	for rows.Next() {
		var currStr string
		err = rows.Scan(&manwhaPage.id, &manwhaPage.pageNumber, &currStr)
		if err != nil {
			log.Warn("cannot scan url %s", err.Error())
			continue
		}
		urls = append(urls, currStr)
	}
	manwhaPage.ImageUrls = urls

	return manwhaPage, nil
}
