package goascii

const (
	//26	1a	SUB / EOF	Substitute / End Of File	文字置換 / ファイル終端
	//27	1b	ESC	Escape	エスケープ(特殊文字開始)
	//28	1c	FS	File Separator	ファイル区切り <--- awk の SUBSEP
	//29	1d	GS	Group Separator	グループ区切り
	//30	1e	RS	Record Separator	レコード区切り
	//31	1f	US	Unit Separator	ユニット区切り
	EOF = rune(26 + iota)
	ESC
	FileSeparator
	GroupSeparator
	RecordSeparator
	UnitSeparator
)
