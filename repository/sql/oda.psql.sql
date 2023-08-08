DROP TABLE IF EXISTS Afstemning;

CREATE TABLE IF NOT EXISTS Afstemning  (
	id INTEGER PRIMARY KEY,
	nummer INTEGER,
	konklusion TEXT NOT NULL,
	vedtaget BOOLEAN,
	kommentar TEXT NOT NULL,
	mødeid INTEGER,
	typeid INTEGER,
	sagstrinid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Afstemningstype;

CREATE TABLE IF NOT EXISTS Afstemningstype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Aktør;

CREATE TABLE IF NOT EXISTS Aktør  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER,
	gruppenavnkort TEXT NOT NULL,
	navn TEXT NOT NULL,
	fornavn TEXT NOT NULL,
	efternavn TEXT NOT NULL,
	biografi TEXT NOT NULL,
	periodeid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP,
	startdato TIMESTAMP NOT NULL,
	slutdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS AktørAktør;

CREATE TABLE IF NOT EXISTS AktørAktør  (
	id INTEGER PRIMARY KEY,
	fraaktørid INTEGER,
	tilaktørid INTEGER,
	startdato TIMESTAMP NOT NULL,
	slutdato TIMESTAMP NOT NULL,
	opdateringsdato TIMESTAMP,
	rolleid INTEGER
);

DROP TABLE IF EXISTS AktørAktørRolle;

CREATE TABLE IF NOT EXISTS AktørAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Aktørtype;

CREATE TABLE IF NOT EXISTS Aktørtype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Dagsordenspunkt;

CREATE TABLE IF NOT EXISTS Dagsordenspunkt  (
	id INTEGER PRIMARY KEY,
	kørebemærkning TEXT NOT NULL,
	titel TEXT NOT NULL,
	kommentar TEXT NOT NULL,
	nummer TEXT NOT NULL,
	forhandlingskode TEXT NOT NULL,
	forhandling TEXT NOT NULL,
	superid INTEGER NOT NULL,
	sagstrinid INTEGER NOT NULL,
	mødeid INTEGER,
	offentlighedskode TEXT,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS DagsordenspunktDokument;

CREATE TABLE IF NOT EXISTS DagsordenspunktDokument  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER,
	dagsordenspunktid INTEGER,
	note TEXT NOT NULL,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS DagsordenspunktSag;

CREATE TABLE IF NOT EXISTS DagsordenspunktSag  (
	id INTEGER PRIMARY KEY,
	dagsordenspunktid INTEGER,
	sagid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Dokument;

CREATE TABLE IF NOT EXISTS Dokument  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER,
	kategoriid INTEGER,
	statusid INTEGER,
	offentlighedskode TEXT,
	titel TEXT,
	dato TIMESTAMP,
	modtagelsesdato TIMESTAMP NOT NULL,
	frigivelsesdato TIMESTAMP NOT NULL,
	paragraf TEXT NOT NULL,
	paragrafnummer TEXT NOT NULL,
	spørgsmålsordlyd TEXT NOT NULL,
	spørgsmålstitel TEXT NOT NULL,
	spørgsmålsid INTEGER NOT NULL,
	procedurenummer TEXT NOT NULL,
	grundnotatstatus TEXT NOT NULL,
	dagsordenudgavenummer SMALLINT NOT NULL,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS DokumentAktør;

CREATE TABLE IF NOT EXISTS DokumentAktør  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER,
	aktørid INTEGER,
	opdateringsdato TIMESTAMP,
	rolleid INTEGER
);

DROP TABLE IF EXISTS DokumentAktørRolle;

CREATE TABLE IF NOT EXISTS DokumentAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Dokumentkategori;

CREATE TABLE IF NOT EXISTS Dokumentkategori  (
	id INTEGER PRIMARY KEY,
	kategori TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Dokumentstatus;

CREATE TABLE IF NOT EXISTS Dokumentstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Dokumenttype;

CREATE TABLE IF NOT EXISTS Dokumenttype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Emneord;

CREATE TABLE IF NOT EXISTS Emneord  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER,
	emneord TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS EmneordDokument;

CREATE TABLE IF NOT EXISTS EmneordDokument  (
	id INTEGER PRIMARY KEY,
	emneordid INTEGER,
	dokumentid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS EmneordSag;

CREATE TABLE IF NOT EXISTS EmneordSag  (
	id INTEGER PRIMARY KEY,
	emneordid INTEGER,
	sagid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Emneordstype;

CREATE TABLE IF NOT EXISTS Emneordstype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS EntitetBeskrivelse;

CREATE TABLE IF NOT EXISTS EntitetBeskrivelse  (
	id INTEGER PRIMARY KEY,
	entitetnavn TEXT NOT NULL,
	beskrivelse TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Fil;

CREATE TABLE IF NOT EXISTS Fil  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER,
	titel TEXT NOT NULL,
	versionsdato TIMESTAMP,
	filurl TEXT,
	opdateringsdato TIMESTAMP,
	variantkode TEXT,
	format TEXT
);

DROP TABLE IF EXISTS KolloneBeskrivelse;

CREATE TABLE IF NOT EXISTS KolloneBeskrivelse  (
	id INTEGER PRIMARY KEY,
	entitetnavn TEXT NOT NULL,
	kollonenavn TEXT NOT NULL,
	beskrivelse TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Møde;

CREATE TABLE IF NOT EXISTS Møde  (
	id INTEGER PRIMARY KEY,
	titel TEXT,
	lokale TEXT NOT NULL,
	nummer TEXT NOT NULL,
	dagsordenurl TEXT NOT NULL,
	starttidsbemærkning TEXT NOT NULL,
	offentlighedskode TEXT,
	dato TIMESTAMP NOT NULL,
	statusid INTEGER NOT NULL,
	typeid INTEGER NOT NULL,
	periodeid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS MødeAktør;

CREATE TABLE IF NOT EXISTS MødeAktør  (
	id INTEGER PRIMARY KEY,
	mødeid INTEGER,
	aktørid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Mødestatus;

CREATE TABLE IF NOT EXISTS Mødestatus  (
	id INTEGER PRIMARY KEY,
	status TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Mødetype;

CREATE TABLE IF NOT EXISTS Mødetype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Omtryk;

CREATE TABLE IF NOT EXISTS Omtryk  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER,
	dato TIMESTAMP NOT NULL,
	begrundelse TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Periode;

CREATE TABLE IF NOT EXISTS Periode  (
	id INTEGER PRIMARY KEY,
	startdato TIMESTAMP,
	slutdato TIMESTAMP,
	type TEXT,
	kode TEXT,
	titel TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Sag;

CREATE TABLE IF NOT EXISTS Sag  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER,
	kategoriid INTEGER NOT NULL,
	statusid INTEGER,
	titel TEXT,
	titelkort TEXT,
	offentlighedskode TEXT,
	nummer TEXT NOT NULL,
	nummerprefix TEXT NOT NULL,
	nummernumerisk TEXT NOT NULL,
	nummerpostfix TEXT NOT NULL,
	resume TEXT NOT NULL,
	afstemningskonklusion TEXT NOT NULL,
	periodeid INTEGER,
	afgørelsesresultatkode TEXT NOT NULL,
	baggrundsmateriale TEXT NOT NULL,
	opdateringsdato TIMESTAMP,
	statsbudgetsag BOOLEAN NOT NULL,
	begrundelse TEXT NOT NULL,
	paragrafnummer INTEGER NOT NULL,
	paragraf TEXT NOT NULL,
	afgørelsesdato TIMESTAMP NOT NULL,
	afgørelse TEXT NOT NULL,
	rådsmødedato TIMESTAMP NOT NULL,
	lovnummer TEXT NOT NULL,
	lovnummerdato TIMESTAMP NOT NULL,
	retsinformationsurl TEXT NOT NULL,
	fremsatundersagid INTEGER NOT NULL,
	deltundersagid INTEGER NOT NULL
);

DROP TABLE IF EXISTS SagAktør;

CREATE TABLE IF NOT EXISTS SagAktør  (
	id INTEGER PRIMARY KEY,
	aktørid INTEGER,
	sagid INTEGER,
	opdateringsdato TIMESTAMP,
	rolleid INTEGER
);

DROP TABLE IF EXISTS SagAktørRolle;

CREATE TABLE IF NOT EXISTS SagAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS SagDokument;

CREATE TABLE IF NOT EXISTS SagDokument  (
	id INTEGER PRIMARY KEY,
	sagid INTEGER,
	dokumentid INTEGER,
	bilagsnummer TEXT NOT NULL,
	frigivelsesdato TIMESTAMP NOT NULL,
	opdateringsdato TIMESTAMP,
	rolleid INTEGER
);

DROP TABLE IF EXISTS SagDokumentRolle;

CREATE TABLE IF NOT EXISTS SagDokumentRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Sagskategori;

CREATE TABLE IF NOT EXISTS Sagskategori  (
	id INTEGER PRIMARY KEY,
	kategori TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Sagsstatus;

CREATE TABLE IF NOT EXISTS Sagsstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Sagstrin;

CREATE TABLE IF NOT EXISTS Sagstrin  (
	id INTEGER PRIMARY KEY,
	titel TEXT,
	dato TIMESTAMP NOT NULL,
	sagid INTEGER,
	typeid INTEGER,
	folketingstidendeurl TEXT NOT NULL,
	folketingstidende TEXT NOT NULL,
	folketingstidendesidenummer TEXT NOT NULL,
	statusid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS SagstrinAktør;

CREATE TABLE IF NOT EXISTS SagstrinAktør  (
	id INTEGER PRIMARY KEY,
	sagstrinid INTEGER,
	aktørid INTEGER,
	opdateringsdato TIMESTAMP,
	rolleid INTEGER
);

DROP TABLE IF EXISTS SagstrinAktørRolle;

CREATE TABLE IF NOT EXISTS SagstrinAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS SagstrinDokument;

CREATE TABLE IF NOT EXISTS SagstrinDokument  (
	id INTEGER PRIMARY KEY,
	sagstrinid INTEGER,
	dokumentid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Sagstrinsstatus;

CREATE TABLE IF NOT EXISTS Sagstrinsstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Sagstrinstype;

CREATE TABLE IF NOT EXISTS Sagstrinstype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Sagstype;

CREATE TABLE IF NOT EXISTS Sagstype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Sambehandlinger;

CREATE TABLE IF NOT EXISTS Sambehandlinger  (
	id INTEGER PRIMARY KEY,
	førstesagstrinid INTEGER,
	andetsagstrinid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Stemme;

CREATE TABLE IF NOT EXISTS Stemme  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER NOT NULL,
	afstemningid INTEGER,
	aktørid INTEGER,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Stemmetype;

CREATE TABLE IF NOT EXISTS Stemmetype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TIMESTAMP
);

