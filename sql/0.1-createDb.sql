DROP TABLE IF EXISTS idea;
CREATE TABLE idea (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    addtime TIMESTAMP ,
    changetime TIMESTAMP  ,
    idea TEXT,
    status INTEGER
);