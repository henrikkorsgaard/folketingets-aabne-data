/* SQL CREATE FILE */

DROP TABLE IF EXISTS sag;

CREATE TABLE IF NOT EXISTS sag (
    id INTEGER PRIMARY KEY,
    type_id INTEGER NOT NULL,
    kategori_id INTEGER,
    status_id INTEGER NOT NULL,
    titel TEXT NOT NULL,
    titel_kort TEXT NOT NULL,
    offentlighedskode TEXT NOT NULL,
    nummer TEXT,
    nummer_prefix TEXT,
    nummer_nummerisk TEXT,
    nummer_postfix TEXT,
    resume TEXT
);