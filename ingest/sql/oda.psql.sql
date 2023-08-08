DROP TABLE IF EXISTS Afstemning;

CREATE TABLE IF NOT EXISTS Afstemning  (
	id INTEGER PRIMARY KEY,
	nummer INTEGER NOT NULL,
	konklusion TEXT,
	vedtaget BOOLEAN NOT NULL,
	kommentar TEXT,
	mødeid INTEGER NOT NULL,
	typeid INTEGER NOT NULL,
	sagstrinid INTEGER,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Afstemningstype;

CREATE TABLE IF NOT EXISTS Afstemningstype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Aktør;

CREATE TABLE IF NOT EXISTS Aktør  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER NOT NULL,
	gruppenavnkort TEXT,
	navn TEXT,
	fornavn TEXT,
	efternavn TEXT,
	biografi TEXT,
	periodeid INTEGER,
	opdateringsdato TIMESTAMP NOT NULL,
	startdato TIMESTAMP,
	slutdato TIMESTAMP
);

DROP TABLE IF EXISTS AktørAktør;

CREATE TABLE IF NOT EXISTS AktørAktør  (
	id INTEGER PRIMARY KEY,
	fraaktørid INTEGER NOT NULL,
	tilaktørid INTEGER NOT NULL,
	startdato TIMESTAMP,
	slutdato TIMESTAMP,
	opdateringsdato TIMESTAMP NOT NULL,
	rolleid INTEGER NOT NULL
);

DROP TABLE IF EXISTS AktørAktørRolle;

CREATE TABLE IF NOT EXISTS AktørAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Aktørtype;

CREATE TABLE IF NOT EXISTS Aktørtype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Dagsordenspunkt;

CREATE TABLE IF NOT EXISTS Dagsordenspunkt  (
	id INTEGER PRIMARY KEY,
	kørebemærkning TEXT,
	titel TEXT,
	kommentar TEXT,
	nummer TEXT,
	forhandlingskode TEXT,
	forhandling TEXT,
	superid INTEGER,
	sagstrinid INTEGER,
	mødeid INTEGER NOT NULL,
	offentlighedskode TEXT NOT NULL,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS DagsordenspunktDokument;

CREATE TABLE IF NOT EXISTS DagsordenspunktDokument  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER NOT NULL,
	dagsordenspunktid INTEGER NOT NULL,
	note TEXT,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS DagsordenspunktSag;

CREATE TABLE IF NOT EXISTS DagsordenspunktSag  (
	id INTEGER PRIMARY KEY,
	dagsordenspunktid INTEGER NOT NULL,
	sagid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Dokument;

CREATE TABLE IF NOT EXISTS Dokument  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER NOT NULL,
	kategoriid INTEGER NOT NULL,
	statusid INTEGER NOT NULL,
	offentlighedskode TEXT NOT NULL,
	titel TEXT NOT NULL,
	dato TIMESTAMP NOT NULL,
	modtagelsesdato TIMESTAMP,
	frigivelsesdato TIMESTAMP,
	paragraf TEXT,
	paragrafnummer TEXT,
	spørgsmålsordlyd TEXT,
	spørgsmålstitel TEXT,
	spørgsmålsid INTEGER,
	procedurenummer TEXT,
	grundnotatstatus TEXT,
	dagsordenudgavenummer SMALLINT,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS DokumentAktør;

CREATE TABLE IF NOT EXISTS DokumentAktør  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER NOT NULL,
	aktørid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL,
	rolleid INTEGER NOT NULL
);

DROP TABLE IF EXISTS DokumentAktørRolle;

CREATE TABLE IF NOT EXISTS DokumentAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Dokumentkategori;

CREATE TABLE IF NOT EXISTS Dokumentkategori  (
	id INTEGER PRIMARY KEY,
	kategori TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Dokumentstatus;

CREATE TABLE IF NOT EXISTS Dokumentstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Dokumenttype;

CREATE TABLE IF NOT EXISTS Dokumenttype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Emneord;

CREATE TABLE IF NOT EXISTS Emneord  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER NOT NULL,
	emneord TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS EmneordDokument;

CREATE TABLE IF NOT EXISTS EmneordDokument  (
	id INTEGER PRIMARY KEY,
	emneordid INTEGER NOT NULL,
	dokumentid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS EmneordSag;

CREATE TABLE IF NOT EXISTS EmneordSag  (
	id INTEGER PRIMARY KEY,
	emneordid INTEGER NOT NULL,
	sagid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Emneordstype;

CREATE TABLE IF NOT EXISTS Emneordstype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS EntitetBeskrivelse;

CREATE TABLE IF NOT EXISTS EntitetBeskrivelse  (
	id INTEGER PRIMARY KEY,
	entitetnavn TEXT,
	beskrivelse TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Fil;

CREATE TABLE IF NOT EXISTS Fil  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER NOT NULL,
	titel TEXT,
	versionsdato TIMESTAMP NOT NULL,
	filurl TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL,
	variantkode TEXT NOT NULL,
	format TEXT NOT NULL
);

DROP TABLE IF EXISTS KolloneBeskrivelse;

CREATE TABLE IF NOT EXISTS KolloneBeskrivelse  (
	id INTEGER PRIMARY KEY,
	entitetnavn TEXT,
	kollonenavn TEXT,
	beskrivelse TEXT,
	opdateringsdato TIMESTAMP
);

DROP TABLE IF EXISTS Møde;

CREATE TABLE IF NOT EXISTS Møde  (
	id INTEGER PRIMARY KEY,
	titel TEXT NOT NULL,
	lokale TEXT,
	nummer TEXT,
	dagsordenurl TEXT,
	starttidsbemærkning TEXT,
	offentlighedskode TEXT NOT NULL,
	dato TIMESTAMP,
	statusid INTEGER,
	typeid INTEGER,
	periodeid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS MødeAktør;

CREATE TABLE IF NOT EXISTS MødeAktør  (
	id INTEGER PRIMARY KEY,
	mødeid INTEGER NOT NULL,
	aktørid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Mødestatus;

CREATE TABLE IF NOT EXISTS Mødestatus  (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Mødetype;

CREATE TABLE IF NOT EXISTS Mødetype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Omtryk;

CREATE TABLE IF NOT EXISTS Omtryk  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER NOT NULL,
	dato TIMESTAMP,
	begrundelse TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Periode;

CREATE TABLE IF NOT EXISTS Periode  (
	id INTEGER PRIMARY KEY,
	startdato TIMESTAMP NOT NULL,
	slutdato TIMESTAMP NOT NULL,
	type TEXT NOT NULL,
	kode TEXT NOT NULL,
	titel TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Sag;

CREATE TABLE IF NOT EXISTS Sag  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER NOT NULL,
	kategoriid INTEGER,
	statusid INTEGER NOT NULL,
	titel TEXT NOT NULL,
	titelkort TEXT NOT NULL,
	offentlighedskode TEXT NOT NULL,
	nummer TEXT,
	nummerprefix TEXT,
	nummernumerisk TEXT,
	nummerpostfix TEXT,
	resume TEXT,
	afstemningskonklusion TEXT,
	periodeid INTEGER NOT NULL,
	afgørelsesresultatkode TEXT,
	baggrundsmateriale TEXT,
	opdateringsdato TIMESTAMP NOT NULL,
	statsbudgetsag BOOLEAN,
	begrundelse TEXT,
	paragrafnummer INTEGER,
	paragraf TEXT,
	afgørelsesdato TIMESTAMP,
	afgørelse TEXT,
	rådsmødedato TIMESTAMP,
	lovnummer TEXT,
	lovnummerdato TIMESTAMP,
	retsinformationsurl TEXT,
	fremsatundersagid INTEGER,
	deltundersagid INTEGER
);

DROP TABLE IF EXISTS SagAktør;

CREATE TABLE IF NOT EXISTS SagAktør  (
	id INTEGER PRIMARY KEY,
	aktørid INTEGER NOT NULL,
	sagid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL,
	rolleid INTEGER NOT NULL
);

DROP TABLE IF EXISTS SagAktørRolle;

CREATE TABLE IF NOT EXISTS SagAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS SagDokument;

CREATE TABLE IF NOT EXISTS SagDokument  (
	id INTEGER PRIMARY KEY,
	sagid INTEGER NOT NULL,
	dokumentid INTEGER NOT NULL,
	bilagsnummer TEXT,
	frigivelsesdato TIMESTAMP,
	opdateringsdato TIMESTAMP NOT NULL,
	rolleid INTEGER NOT NULL
);

DROP TABLE IF EXISTS SagDokumentRolle;

CREATE TABLE IF NOT EXISTS SagDokumentRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Sagskategori;

CREATE TABLE IF NOT EXISTS Sagskategori  (
	id INTEGER PRIMARY KEY,
	kategori TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Sagsstatus;

CREATE TABLE IF NOT EXISTS Sagsstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Sagstrin;

CREATE TABLE IF NOT EXISTS Sagstrin  (
	id INTEGER PRIMARY KEY,
	titel TEXT NOT NULL,
	dato TIMESTAMP,
	sagid INTEGER NOT NULL,
	typeid INTEGER NOT NULL,
	folketingstidendeurl TEXT,
	folketingstidende TEXT,
	folketingstidendesidenummer TEXT,
	statusid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS SagstrinAktør;

CREATE TABLE IF NOT EXISTS SagstrinAktør  (
	id INTEGER PRIMARY KEY,
	sagstrinid INTEGER NOT NULL,
	aktørid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL,
	rolleid INTEGER NOT NULL
);

DROP TABLE IF EXISTS SagstrinAktørRolle;

CREATE TABLE IF NOT EXISTS SagstrinAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS SagstrinDokument;

CREATE TABLE IF NOT EXISTS SagstrinDokument  (
	id INTEGER PRIMARY KEY,
	sagstrinid INTEGER NOT NULL,
	dokumentid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Sagstrinsstatus;

CREATE TABLE IF NOT EXISTS Sagstrinsstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Sagstrinstype;

CREATE TABLE IF NOT EXISTS Sagstrinstype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Sagstype;

CREATE TABLE IF NOT EXISTS Sagstype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Sambehandlinger;

CREATE TABLE IF NOT EXISTS Sambehandlinger  (
	id INTEGER PRIMARY KEY,
	førstesagstrinid INTEGER NOT NULL,
	andetsagstrinid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Stemme;

CREATE TABLE IF NOT EXISTS Stemme  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER,
	afstemningid INTEGER NOT NULL,
	aktørid INTEGER NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

DROP TABLE IF EXISTS Stemmetype;

CREATE TABLE IF NOT EXISTS Stemmetype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TIMESTAMP NOT NULL
);

