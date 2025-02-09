package drivers

const (
	ManwhaResourceTable string = "ManwhaResource"
	ManwhaPageTable     string = "ManwhaPage"
	ManwhaPageURLTable  string = "ManwhaPageURL"

	CreateManwhaResourceTable string = `
	CREATE TABLE IF NOT EXISTS ` + ManwhaResourceTable + `  (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		address  TEXT NOT NULL,
		nChapter INT NOT NULL,
		imageURL TEXT NOT NULL,
		driver TEXT
	)
	`

	CreateManwhaPageTable string = `
	CREATE TABLE IF NOT EXISTS ` + ManwhaPageTable + `  (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		pageNumber INTEGER NOT NULL,
		resourceId INTEGER NOT NULL ,
		FOREIGN KEY (resourceId) REFERENCES ` + ManwhaResourceTable + `(id)
	)
	`

	CreateManwhaPageURLTable string = `
	CREATE TABLE IF NOT EXISTS ` + ManwhaPageURLTable + `  (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		imageUrl Text NOT NULL,
		pageId INTEGER NOT NULL ,
		FOREIGN KEY (pageId) REFERENCES ManwhaPage(id)
	)
	`
	InsertManwhaResource string = `
	INSERT INTO ` + ManwhaResourceTable + ` (name,address,nChapter,imageURL,driver)
	VALUES($1,$2,$3,$4,$5) 
	`

	InsertManwhaPage string = `
	INSERT INTO ` + ManwhaPageTable + ` (pageNumber,resourceId)
	VALUES($1,$2) 
	`

	InsertManwhaPageURL string = `
	INSERT INTO ` + ManwhaPageURLTable + ` (imageUrl,pageId)
	VALUES($1,$2) 
	`

	InsertIfNotExistManwhaResource string = `
	INSERT INTO ` + ManwhaResourceTable + ` (name,address,nChapter,imageURL,driver)
	SELECT $1, $2, $3, $4, $5
	WHERE NOT EXISTS(SELECT 1 FROM ` + ManwhaResourceTable + ` WHERE name=$1 AND driver=$5);
	SELECT id FROM ` + ManwhaResourceTable + ` WHERE name=$1 AND driver=$5
	`

	InsertIfNotExistManwhaPage string = `
	INSERT INTO ` + ManwhaPageTable + ` (pageNumber,resourceId)
	SELECT $1, $2
	WHERE NOT EXISTS(SELECT 1 FROM ` + ManwhaPageTable + ` WHERE pageNumber=$1 AND resourceId=$2);
	SELECT id FROM ` + ManwhaPageTable + ` WHERE pageNumber=$1 AND resourceId=$2
	`

	InsertIfNotExistManwhaPageURL string = `
	INSERT INTO ` + ManwhaPageURLTable + ` (imageUrl,pageId)
	SELECT $1, $2
	WHERE NOT EXISTS(SELECT 1 FROM ` + ManwhaPageURLTable + ` WHERE imageUrl=$1 AND pageId=$2);
	SELECT id FROM ` + ManwhaPageURLTable + ` WHERE imageUrl=$1 AND pageId=$2
	`

	SelectManwhaResource string = `
	SELECT * FROM ` + ManwhaResourceTable + `
	`
	SelectManwhaPage string = `
	SELECT * FROM ` + ManwhaPageTable + `
	`
	SelectManwhaPageURL string = `
	SELECT * FROM ` + ManwhaPageURLTable + `
	`
	SelectManwhaResourceById string = `
	SELECT * FROM ` + ManwhaResourceTable + ` WHERE id=$1
	`
	SelectManwhaPageById string = `
	SELECT * FROM  ` + ManwhaPageTable + ` WHERE id=$1
	`
	SelectManwhaPageURLById string = `
	SELECT * FROM ` + ManwhaPageURLTable + ` WHERE id=$1 
	`
	SelectManwhaPageByResourceId string = `
	SELECT * FROM  ` + ManwhaPageTable + ` WHERE resourceId=$1
	`
	SelectManwhaPageURLByPageId string = `
	SELECT * FROM ` + ManwhaPageURLTable + ` WHERE pageId=$1 
	`
	JoinManwhaPage string = `
	SELECT ` + ManwhaPageTable + `.id,ManwhaPage.pageNumber,ManwhaPageURL.imageUrl FROM  ` + ManwhaPageTable + `
	JOIN ` + ManwhaPageURLTable + ` ON ManwhaPage.id=ManwhaPageURL.pageId WHERE ManwhaPage.id=$1
	`
	JoinManwhaPageByResourceId string = `
	SELECT ` + ManwhaPageTable + `.id,ManwhaPage.pageNumber,ManwhaPageURL.imageUrl FROM  ` + ManwhaPageTable + `
	JOIN ` + ManwhaPageURLTable + ` ON ManwhaPage.id=ManwhaPageURL.pageId WHERE ManwhaPage.resourceId=$1
	`
)

type database interface {
	InitDatabase(force bool) error
	AddManwhaResource(resource *ManwhaResource) error
	AddManwhaPage(page *ManwhaPage) error
	GetManwhaResource(id int, collectPages bool) (*ManwhaResource, error)
	GetManwhaPage(id int) (*ManwhaPage, error)
}
