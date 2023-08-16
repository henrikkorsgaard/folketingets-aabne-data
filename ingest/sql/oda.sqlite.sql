CREATE TABLE IF NOT EXISTS Afstemning  (
	id INTEGER PRIMARY KEY,
	nummer INTEGER NOT NULL,
	konklusion TEXT,
	vedtaget INTEGER NOT NULL,
	kommentar TEXT,
	mødeid INTEGER NOT NULL,
	typeid INTEGER NOT NULL,
	sagstrinid INTEGER,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Afstemningstype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Aktør  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER NOT NULL,
	gruppenavnkort TEXT,
	navn TEXT,
	fornavn TEXT,
	efternavn TEXT,
	biografi TEXT,
	periodeid INTEGER,
	opdateringsdato TEXT NOT NULL,
	startdato TEXT,
	slutdato TEXT
);

CREATE TABLE IF NOT EXISTS AktørAktør  (
	id INTEGER PRIMARY KEY,
	fraaktørid INTEGER NOT NULL,
	tilaktørid INTEGER NOT NULL,
	startdato TEXT,
	slutdato TEXT,
	opdateringsdato TEXT NOT NULL,
	rolleid INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS AktørAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Aktørtype  (
	id INTEGER PRIMARY KEY,
	type TEXT,
	opdateringsdato TEXT NOT NULL
);

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
	opdateringsdato TEXT
);

CREATE TABLE IF NOT EXISTS DagsordenspunktDokument  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER NOT NULL,
	dagsordenspunktid INTEGER NOT NULL,
	note TEXT,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS DagsordenspunktSag  (
	id INTEGER PRIMARY KEY,
	dagsordenspunktid INTEGER NOT NULL,
	sagid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Dokument  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER NOT NULL,
	kategoriid INTEGER NOT NULL,
	statusid INTEGER NOT NULL,
	offentlighedskode TEXT NOT NULL,
	titel TEXT NOT NULL,
	dato TEXT NOT NULL,
	modtagelsesdato TEXT,
	frigivelsesdato TEXT,
	paragraf TEXT,
	paragrafnummer TEXT,
	spørgsmålsordlyd TEXT,
	spørgsmålstitel TEXT,
	spørgsmålsid INTEGER,
	procedurenummer TEXT,
	grundnotatstatus TEXT,
	dagsordenudgavenummer INTEGER,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS DokumentAktør  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER NOT NULL,
	aktørid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL,
	rolleid INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS DokumentAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Dokumentkategori  (
	id INTEGER PRIMARY KEY,
	kategori TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Dokumentstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Dokumenttype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Emneord  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER NOT NULL,
	emneord TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS EmneordDokument  (
	id INTEGER PRIMARY KEY,
	emneordid INTEGER NOT NULL,
	dokumentid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS EmneordSag  (
	id INTEGER PRIMARY KEY,
	emneordid INTEGER NOT NULL,
	sagid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Emneordstype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS EntitetBeskrivelse  (
	id INTEGER PRIMARY KEY,
	entitetnavn TEXT,
	beskrivelse TEXT,
	opdateringsdato TEXT
);

CREATE TABLE IF NOT EXISTS Fil  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER NOT NULL,
	titel TEXT,
	versionsdato TEXT NOT NULL,
	filurl TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL,
	variantkode TEXT NOT NULL,
	format TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS KolloneBeskrivelse  (
	id INTEGER PRIMARY KEY,
	entitetnavn TEXT,
	kollonenavn TEXT,
	beskrivelse TEXT,
	opdateringsdato TEXT
);

CREATE TABLE IF NOT EXISTS Møde  (
	id INTEGER PRIMARY KEY,
	titel TEXT NOT NULL,
	lokale TEXT,
	nummer TEXT,
	dagsordenurl TEXT,
	starttidsbemærkning TEXT,
	offentlighedskode TEXT NOT NULL,
	dato TEXT,
	statusid INTEGER,
	typeid INTEGER,
	periodeid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS MødeAktør  (
	id INTEGER PRIMARY KEY,
	mødeid INTEGER NOT NULL,
	aktørid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Mødestatus  (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Mødetype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Omtryk  (
	id INTEGER PRIMARY KEY,
	dokumentid INTEGER NOT NULL,
	dato TEXT,
	begrundelse TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Periode  (
	id INTEGER PRIMARY KEY,
	startdato TEXT NOT NULL,
	slutdato TEXT NOT NULL,
	type TEXT NOT NULL,
	kode TEXT NOT NULL,
	titel TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

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
	opdateringsdato TEXT NOT NULL,
	statsbudgetsag INTEGER,
	begrundelse TEXT,
	paragrafnummer INTEGER,
	paragraf TEXT,
	afgørelsesdato TEXT,
	afgørelse TEXT,
	rådsmødedato TEXT,
	lovnummer TEXT,
	lovnummerdato TEXT,
	retsinformationsurl TEXT,
	fremsatundersagid INTEGER,
	deltundersagid INTEGER
);

CREATE TABLE IF NOT EXISTS SagAktør  (
	id INTEGER PRIMARY KEY,
	aktørid INTEGER NOT NULL,
	sagid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL,
	rolleid INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS SagAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS SagDokument  (
	id INTEGER PRIMARY KEY,
	sagid INTEGER NOT NULL,
	dokumentid INTEGER NOT NULL,
	bilagsnummer TEXT,
	frigivelsesdato TEXT,
	opdateringsdato TEXT NOT NULL,
	rolleid INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS SagDokumentRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Sagskategori  (
	id INTEGER PRIMARY KEY,
	kategori TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Sagsstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Sagstrin  (
	id INTEGER PRIMARY KEY,
	titel TEXT NOT NULL,
	dato TEXT,
	sagid INTEGER NOT NULL,
	typeid INTEGER NOT NULL,
	folketingstidendeurl TEXT,
	folketingstidende TEXT,
	folketingstidendesidenummer TEXT,
	statusid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS SagstrinAktør  (
	id INTEGER PRIMARY KEY,
	sagstrinid INTEGER NOT NULL,
	aktørid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL,
	rolleid INTEGER NOT NULL
);

CREATE TABLE IF NOT EXISTS SagstrinAktørRolle  (
	id INTEGER PRIMARY KEY,
	rolle TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS SagstrinDokument  (
	id INTEGER PRIMARY KEY,
	sagstrinid INTEGER NOT NULL,
	dokumentid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Sagstrinsstatus  (
	id INTEGER PRIMARY KEY,
	status TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Sagstrinstype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Sagstype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Sambehandlinger  (
	id INTEGER PRIMARY KEY,
	førstesagstrinid INTEGER NOT NULL,
	andetsagstrinid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Stemme  (
	id INTEGER PRIMARY KEY,
	typeid INTEGER,
	afstemningid INTEGER NOT NULL,
	aktørid INTEGER NOT NULL,
	opdateringsdato TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS Stemmetype  (
	id INTEGER PRIMARY KEY,
	type TEXT NOT NULL,
	opdateringsdato TEXT NOT NULL
);

