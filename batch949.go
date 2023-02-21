package batch949

import (
	"os/exec"
	"syscall"

	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/transform"
)

type CodePage string

const (
	CHCP_IBM037                  CodePage = "037"   //	IBM EBCDIC US-Canada
	CHCP_IBM437                  CodePage = "437"   //	OEM United States
	CHCP_IBM500                  CodePage = "500"   //	IBM EBCDIC International
	CHCP_ASMO_708                CodePage = "708"   //	Arabic (ASMO 708)
	CHCP_DOS_720                 CodePage = "720"   //	Arabic (Transparent ASMO); Arabic (DOS)
	CHCP_ibm737                  CodePage = "737"   //	OEM Greek (formerly 437G); Greek (DOS)
	CHCP_ibm775                  CodePage = "775"   //	OEM Baltic; Baltic (DOS)
	CHCP_ibm850                  CodePage = "850"   //	OEM Multilingual Latin 1; Western European (DOS)
	CHCP_ibm852                  CodePage = "852"   //	OEM Latin 2; Central European (DOS)
	CHCP_IBM855                  CodePage = "855"   //	OEM Cyrillic (primarily Russian)
	CHCP_ibm857                  CodePage = "857"   //	OEM Turkish; Turkish (DOS)
	CHCP_IBM00858                CodePage = "858"   //	OEM Multilingual Latin 1 + Euro symbol
	CHCP_IBM860                  CodePage = "860"   //	OEM Portuguese; Portuguese (DOS)
	CHCP_ibm861                  CodePage = "861"   //	OEM Icelandic; Icelandic (DOS)
	CHCP_DOS_862                 CodePage = "862"   //	OEM Hebrew; Hebrew (DOS)
	CHCP_IBM863                  CodePage = "863"   //	OEM French Canadian; French Canadian (DOS)
	CHCP_IBM864                  CodePage = "864"   //	OEM Arabic; Arabic (864)
	CHCP_IBM865                  CodePage = "865"   //	OEM Nordic; Nordic (DOS)
	CHCP_cp866                   CodePage = "866"   //	OEM Russian; Cyrillic (DOS)
	CHCP_ibm869                  CodePage = "869"   //	OEM Modern Greek; Greek, Modern (DOS)
	CHCP_IBM870                  CodePage = "870"   //	IBM EBCDIC Multilingual/ROECE (Latin 2); IBM EBCDIC Multilingual Latin 2
	CHCP_windows_874             CodePage = "874"   //	Thai (Windows)
	CHCP_cp875                   CodePage = "875"   //	IBM EBCDIC Greek Modern
	CHCP_shift_jis               CodePage = "932"   //	ANSI/OEM Japanese; Japanese (Shift-JIS)
	CHCP_gb2312                  CodePage = "936"   //	ANSI/OEM Simplified Chinese (PRC, Singapore); Chinese Simplified (GB2312)
	CHCP_ks_c_5601_1987          CodePage = "949"   //	ANSI/OEM Korean (Unified Hangul Code)
	CHCP_big5                    CodePage = "950"   //	ANSI/OEM Traditional Chinese (Taiwan; Hong Kong SAR, PRC); Chinese Traditional (Big5)
	CHCP_IBM1026                 CodePage = "1026"  //	IBM EBCDIC Turkish (Latin 5)
	CHCP_IBM01047                CodePage = "1047"  //	IBM EBCDIC Latin 1/Open System
	CHCP_IBM01140                CodePage = "1140"  //	IBM EBCDIC US-Canada (037 + Euro symbol); IBM EBCDIC (US-Canada-Euro)
	CHCP_IBM01141                CodePage = "1141"  //	IBM EBCDIC Germany (20273 + Euro symbol); IBM EBCDIC (Germany-Euro)
	CHCP_IBM01142                CodePage = "1142"  //	IBM EBCDIC Denmark-Norway (20277 + Euro symbol); IBM EBCDIC (Denmark-Norway-Euro)
	CHCP_IBM01143                CodePage = "1143"  //	IBM EBCDIC Finland-Sweden (20278 + Euro symbol); IBM EBCDIC (Finland-Sweden-Euro)
	CHCP_IBM01144                CodePage = "1144"  //	IBM EBCDIC Italy (20280 + Euro symbol); IBM EBCDIC (Italy-Euro)
	CHCP_IBM01145                CodePage = "1145"  //	IBM EBCDIC Latin America-Spain (20284 + Euro symbol); IBM EBCDIC (Spain-Euro)
	CHCP_IBM01146                CodePage = "1146"  //	IBM EBCDIC United Kingdom (20285 + Euro symbol); IBM EBCDIC (UK-Euro)
	CHCP_IBM01147                CodePage = "1147"  //	IBM EBCDIC France (20297 + Euro symbol); IBM EBCDIC (France-Euro)
	CHCP_IBM01148                CodePage = "1148"  //	IBM EBCDIC International (500 + Euro symbol); IBM EBCDIC (International-Euro)
	CHCP_IBM01149                CodePage = "1149"  //	IBM EBCDIC Icelandic (20871 + Euro symbol); IBM EBCDIC (Icelandic-Euro)
	CHCP_windows_1250            CodePage = "1250"  //	ANSI Central European; Central European (Windows)
	CHCP_windows_1251            CodePage = "1251"  //	ANSI Cyrillic; Cyrillic (Windows)
	CHCP_windows_1252            CodePage = "1252"  //	ANSI Latin 1; Western European (Windows)
	CHCP_windows_1253            CodePage = "1253"  //	ANSI Greek; Greek (Windows)
	CHCP_windows_1254            CodePage = "1254"  //	ANSI Turkish; Turkish (Windows)
	CHCP_windows_1255            CodePage = "1255"  //	ANSI Hebrew; Hebrew (Windows)
	CHCP_windows_1256            CodePage = "1256"  //	ANSI Arabic; Arabic (Windows)
	CHCP_windows_1257            CodePage = "1257"  //	ANSI Baltic; Baltic (Windows)
	CHCP_windows_1258            CodePage = "1258"  //	ANSI/OEM Vietnamese; Vietnamese (Windows)
	CHCP_Johab                   CodePage = "1361"  //	Korean (Johab)
	CHCP_macintosh               CodePage = "10000" //	MAC Roman; Western European (Mac)
	CHCP_x_mac_japanese          CodePage = "10001" //	Japanese (Mac)
	CHCP_x_mac_chinesetrad       CodePage = "10002" //	MAC Traditional Chinese (Big5); Chinese Traditional (Mac)
	CHCP_x_mac_korean            CodePage = "10003" //	Korean (Mac)
	CHCP_x_mac_arabic            CodePage = "10004" //	Arabic (Mac)
	CHCP_x_mac_hebrew            CodePage = "10005" //	Hebrew (Mac)
	CHCP_x_mac_greek             CodePage = "10006" //	Greek (Mac)
	CHCP_x_mac_cyrillic          CodePage = "10007" //	Cyrillic (Mac)
	CHCP_x_mac_chinesesimp       CodePage = "10008" //	MAC Simplified Chinese (GB 2312); Chinese Simplified (Mac)
	CHCP_x_mac_romanian          CodePage = "10010" //	Romanian (Mac)
	CHCP_x_mac_ukrainian         CodePage = "10017" //	Ukrainian (Mac)
	CHCP_x_mac_thai              CodePage = "10021" //	Thai (Mac)
	CHCP_x_mac_ce                CodePage = "10029" //	MAC Latin 2; Central European (Mac)
	CHCP_x_mac_icelandic         CodePage = "10079" //	Icelandic (Mac)
	CHCP_x_mac_turkish           CodePage = "10081" //	Turkish (Mac)
	CHCP_x_mac_croatian          CodePage = "10082" //	Croatian (Mac)
	CHCP_x_Chinese_CNS           CodePage = "20000" //	CNS Taiwan; Chinese Traditional (CNS)
	CHCP_x_cp20001               CodePage = "20001" //	TCA Taiwan
	CHCP_x_Chinese_Eten          CodePage = "20002" //	Eten Taiwan; Chinese Traditional (Eten)
	CHCP_x_cp20003               CodePage = "20003" //	IBM5550 Taiwan
	CHCP_x_cp20004               CodePage = "20004" //	TeleText Taiwan
	CHCP_x_cp20005               CodePage = "20005" //	Wang Taiwan
	CHCP_x_IA5                   CodePage = "20105" //	IA5 (IRV International Alphabet No. 5, 7_bit); Western European (IA5)
	CHCP_x_IA5_German            CodePage = "20106" //	IA5 German (7_bit)
	CHCP_x_IA5_Swedish           CodePage = "20107" //	IA5 Swedish (7_bit)
	CHCP_x_IA5_Norwegian         CodePage = "20108" //	IA5 Norwegian (7_bit)
	CHCP_us_ascii                CodePage = "20127" //	US_ASCII (7_bit)
	CHCP_x_cp20269               CodePage = "20269" //	ISO 6937 Non_Spacing Accent
	CHCP_IBM273                  CodePage = "20273" //	IBM EBCDIC Germany
	CHCP_IBM277                  CodePage = "20277" //	IBM EBCDIC Denmark_Norway
	CHCP_IBM278                  CodePage = "20278" //	IBM EBCDIC Finland_Sweden
	CHCP_IBM280                  CodePage = "20280" //	IBM EBCDIC Italy
	CHCP_IBM284                  CodePage = "20284" //	IBM EBCDIC Latin America_Spain
	CHCP_IBM285                  CodePage = "20285" //	IBM EBCDIC United Kingdom
	CHCP_IBM290                  CodePage = "20290" //	IBM EBCDIC Japanese Katakana Extended
	CHCP_IBM297                  CodePage = "20297" //	IBM EBCDIC France
	CHCP_IBM420                  CodePage = "20420" //	IBM EBCDIC Arabic
	CHCP_IBM423                  CodePage = "20423" //	IBM EBCDIC Greek
	CHCP_IBM424                  CodePage = "20424" //	IBM EBCDIC Hebrew
	CHCP_x_EBCDIC_KoreanExtended CodePage = "20833" //	IBM EBCDIC Korean Extended
	CHCP_IBM_Thai                CodePage = "20838" //	IBM EBCDIC Thai
	CHCP_koi8_r                  CodePage = "20866" //	Russian (KOI8_R); Cyrillic (KOI8_R)
	CHCP_IBM871                  CodePage = "20871" //	IBM EBCDIC Icelandic
	CHCP_IBM880                  CodePage = "20880" //	IBM EBCDIC Cyrillic Russian
	CHCP_IBM905                  CodePage = "20905" //	IBM EBCDIC Turkish
	CHCP_IBM00924                CodePage = "20924" //	IBM EBCDIC Latin 1/Open System (1047 + Euro symbol)
	CHCP_EUC_JP                  CodePage = "20932" //	Japanese (JIS 0208_1990 and 0212_1990)
	CHCP_x_cp20936               CodePage = "20936" //	Simplified Chinese (GB2312); Chinese Simplified (GB2312_80)
	CHCP_x_cp20949               CodePage = "20949" //	Korean Wansung
	CHCP_cp1025                  CodePage = "21025" //	IBM EBCDIC Cyrillic Serbian_Bulgarian
	CHCP_koi8_u                  CodePage = "21866" //	Ukrainian (KOI8_U); Cyrillic (KOI8_U)
	CHCP_iso_8859_1              CodePage = "28591" //	ISO 8859_1 Latin 1; Western European (ISO)
	CHCP_iso_8859_2              CodePage = "28592" //	ISO 8859_2 Central European; Central European (ISO)
	CHCP_iso_8859_3              CodePage = "28593" //	ISO 8859_3 Latin 3
	CHCP_iso_8859_4              CodePage = "28594" //	ISO 8859_4 Baltic
	CHCP_iso_8859_5              CodePage = "28595" //	ISO 8859_5 Cyrillic
	CHCP_iso_8859_6              CodePage = "28596" //	ISO 8859_6 Arabic
	CHCP_iso_8859_7              CodePage = "28597" //	ISO 8859_7 Greek
	CHCP_iso_8859_8              CodePage = "28598" //	ISO 8859_8 Hebrew; Hebrew (ISO_Visual)
	CHCP_iso_8859_9              CodePage = "28599" //	ISO 8859_9 Turkish
	CHCP_iso_8859_13             CodePage = "28603" //	ISO 8859_13 Estonian
	CHCP_iso_8859_15             CodePage = "28605" //	ISO 8859_15 Latin 9
	CHCP_iso_8859_8_i            CodePage = "38598" //	ISO 8859_8 Hebrew; Hebrew (ISO_Logical)
	CHCP_iso_2022_jp_50220       CodePage = "50220" //	ISO 2022 Japanese with no halfwidth Katakana; Japanese (JIS)
	CHCP_csISO2022JP             CodePage = "50221" //	ISO 2022 Japanese with halfwidth Katakana; Japanese (JIS_Allow 1 byte Kana)
	CHCP_iso_2022_jp             CodePage = "50222" //	ISO 2022 Japanese JIS X 0201_1989; Japanese (JIS_Allow 1 byte Kana _ SO/SI)
	CHCP_iso_2022_kr             CodePage = "50225" //	ISO 2022 Korean
	CHCP_x_cp50227               CodePage = "50227" //	ISO 2022 Simplified Chinese; Chinese Simplified (ISO 2022)
	CHCP_50229                   CodePage = "50229" //	ISO 2022 Traditional Chinese
	CHCP_euc_kr                  CodePage = "51949" //	EUC Korean
	CHCP_hz_gb_2312              CodePage = "52936" //	HZ_GB2312 Simplified Chinese; Chinese Simplified (HZ)
	CHCP_GB18030                 CodePage = "54936" //	Windows XP and later: GB18030 Simplified Chinese (4 byte); Chinese Simplified (GB18030)
	CHCP_x_iscii_de              CodePage = "57002" //	ISCII Devanagari
	CHCP_x_iscii_be              CodePage = "57003" //	ISCII Bangla
	CHCP_x_iscii_ta              CodePage = "57004" //	ISCII Tamil
	CHCP_x_iscii_te              CodePage = "57005" //	ISCII Telugu
	CHCP_x_iscii_as              CodePage = "57006" //	ISCII Assamese
	CHCP_x_iscii_or              CodePage = "57007" //	ISCII Odia
	CHCP_x_iscii_ka              CodePage = "57008" //	ISCII Kannada
	CHCP_x_iscii_ma              CodePage = "57009" //	ISCII Malayalam
	CHCP_x_iscii_gu              CodePage = "57010" //	ISCII Gujarati
	CHCP_x_iscii_pa              CodePage = "57011" //	ISCII Punjabi
	CHCP_utf_7                   CodePage = "65000" //	Unicode (UTF_7)
	CHCP_utf_8                   CodePage = "65001" //	Unicode (UTF_8)
)

var (
	fixedCodePage CodePage
)

func init() {
	fixedCodePage = CHCP_ks_c_5601_1987 // chcp 949 (korean)
}

/**
*	SetCodePage
**/
func SetCodePage(cp CodePage) {
	fixedCodePage = cp
}

/**
*	buildCmd
**/
func buildCmd(execCmd string, args ...string) *exec.Cmd {
	na := []string{"/C", "chcp", string(fixedCodePage), "&", execCmd}
	na = append(na, args...)

	cmd := exec.Command("cmd.exe", na...)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd
}

/**
*	Output
**/
func Output(execCmd string, args ...string) (string, error) {
	cmd := buildCmd(execCmd, args...)

	buf, _ := cmd.CombinedOutput()

	ub, _, ubErr := transform.Bytes(korean.EUCKR.NewDecoder(), buf)

	return string(ub), ubErr
}

/**
*	Run
**/
func Run(execCmd string, args ...string) error {
	cmd := buildCmd(execCmd, args...)
	return cmd.Run()
}

/**
*	Start
**/
func Start(execCmd string, args ...string) error {
	cmd := buildCmd(execCmd, args...)
	return cmd.Start()
}
