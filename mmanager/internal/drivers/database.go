package drivers

const (
	CreateManwhaResourceTable string = `
	CREATE TABLE IF NOT EXISTS ManwhaResource  (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		address  TEXT NOT NULL,
		nChapter INT NOT NULL,
		imageURL TEXT NOT NULL,
		driver TEXT
	)
	`

	CreateManwhaPageTable string = `
	CREATE TABLE IF NOT EXISTS ManwhaPage  (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		pageNumber INTEGER NOT NULL,
		resourceId INTEGER NOT NULL ,
		FOREIGN KEY (resourceId) REFERENCES ManwhaResource(id)
	)
	`

	CreateManwhaPageURLTable string = `
	CREATE TABLE IF NOT EXISTS ManwhaPageURL  (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		imageUrl Text NOT NULL,
		pageId INTEGER NOT NULL ,
		FOREIGN KEY (pageId) REFERENCES ManwhaPage(id)
	)
	`

	InsertManwhaResource string = `
	INSERT INTO ManwhaResource (name,address,nChapter,imageURL,driver)
	VALUES($1,$2,$3,$4,$5) RETURNING id
	`

	InsertManwhaPage string = `
	INSERT INTO ManwhaPage (pageNumber,resourceId)
	VALUES($1,$2) RETURNING id
	`

	InsertManwhaPageURL string = `
	INSERT INTO ManwhaPageURL (imageUrl,pageId)
	VALUES($1,$2) RETURNING id
	`
	SelectManwhaResource string = `
	SELECT * FROM ManwhaResource
	`
	SelectManwhaPage string = `
	SELECT * FROM ManwhaPage
	`
	SelectManwhaPageURL string = `
	SELECT * FROM ManwhaPageURL
	`
	SelectManwhaResourceById string = `
	SELECT * FROM ManwhaResource WHERE id=$1
	`
	SelectManwhaPageById string = `
	SELECT * FROM ManwhaPage WHERE id=$1
	`
	SelectManwhaPageURLById string = `
	SELECT * FROM ManwhaPageURL WHERE id=$1 
	`
	SelectManwhaPageByResourceId string = `
	SELECT * FROM ManwhaPage WHERE resourceId=$1
	`
	SelectManwhaPageURLByPageId string = `
	SELECT * FROM ManwhaPageURL WHERE pageId=$1 
	`
	JoinManwhaPage string = `
	SELECT ManwhaPage.id,ManwhaPage.pageNumber,ManwhaPageURL.imageUrl FROM ManwhaPage
	JOIN ManwhaPageURL ON ManwhaPage.id=ManwhaPageURL.pageId WHERE ManwhaPage.id=$1
	`
	JoinManwhaPageByResourceId string = `
	SELECT ManwhaPage.id,ManwhaPage.pageNumber,ManwhaPageURL.imageUrl FROM ManwhaPage
	JOIN ManwhaPageURL ON ManwhaPage.id=ManwhaPageURL.pageId WHERE ManwhaPage.resourceId=$1
	`
)

type database interface {
	InitDatabase(force bool) error
	AddManwhaResource(resource *ManwhaResource) error
	AddManwhaPage(page *ManwhaPage) error
	GetManwhaResource(id int, collectPages bool) (*ManwhaResource, error)
	GetManwhaPage(id int) (*ManwhaPage, error)
}
