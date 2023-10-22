package stats

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"
)

var playersList = []string{"010Gambler", "0kdoei", "1313Tom1313", "1961Gambler", "1973p0kerpro", "19vladimir831", "1mcsiveeni", "21josemacedo21", "232323", "262329d3c08541d79b4f00f488f93d", "2815708", "35yakarca35", "38750honda", "3betChauvvvvve", "404kate", "4419706", "581011181935", "666Mmt", "66Luuseri", "777queenhappy777", "77shubidua", "7d2h", "8527226865", "96Ntguilty", "A1arak", "AA19841", "AAIdas", "AAKKEU", "AAbasAA", "ABI1kk", "ACEWILSON888", "ACSEMEGE", "ALLINANDHOPE", "ARATATATA", "AbdelA15er", "Acecrakers79", "AcesGrey", "Acid013", "Agent644", "Aigars13", "Aipzer18", "AiratK", "Al7788ex", "Alchertim", "Alcyoneus", "AlexCartel", "Alfa1870", "AlfaRomeo211", "AlfredoDeLaVega", "Almani37", "Aloma", "AmarilloSlimFisu", "Ambervays1", "AndrewLloyd", "Anfang", "ArhPro666", "Arrowxl", "Arruolato", "Arturogatti461", "ArturovichX", "AslanVanacha2", "Attali189", "Attuirmad4", "AwarnessSlave", "Awdowell", "Azaki", "AzevedoMH", "B1llyTalent", "BADASSMOFOOOOO", "BADCHIC999", "BANCIUL1978", "BAYA27", "BEERandFish", "BERICICA36", "BFEdgeguy", "BFKodijack", "BFlynn", "BGUSZTO", "BKNZHN", "BLUEmagic", "BONGOPOKER4ever", "BOTTGSM1", "BackInCCCP1", "BadBoyBrian", "BadFriday", "Bajax83", "Bamboocha35", "Baptist075", "Barbashoe", "Barneyandholly", "Barns54", "Basel4056", "BazookaJoe", "BeGosIV", "Beach6", "BeenRivered1", "Behlings09", "Belle1961", "BennyHendrix", "BenoitBlanc12", "Bienvenido", "Bierpokkell", "BigSidToilet", "BigT1687", "Bingo1700", "BlackStar", "BlakBoxPoker", "BlaueBeere", "Blazing247", "Blondie69", "Bloodmoon", "BlueEssence", "BluffMeNot", "Bogdangready", "BombenHagel", "Boorik1", "Borislavhs", "BruceVeine", "Bruintje67", "Bu789", "BuliVeivari", "Bumm52", "C4rli89", "CARLOSDIEGO71", "CARLOSKILLS4U", "CHEGUEVARA3", "CKOPOBEPHY", "CLAIREBEAR891", "CLAUDIAG69", "CLUSTRET1", "CNML2022", "COREANORICO", "CORK", "CROMER1952", "CZIRJAKSZ17", "Cappel", "CaptTuttle", "Captainkiano", "Carpycall", "Chaefeeli84", "Charlie08", "Charn", "Chiborras", "ChiefPh", "Chino576677", "ChipDouglas1", "Chipaev", "ChoccyCookie", "ChocoBrownie", "Cinneide93", "Claudner", "CleanTable10", "Cloges1", "CogDogy", "Come2Pappa", "Cooboo", "Coolspot", "Corben313", "Cord7777", "CoronaST", "Costa900", "Cougey1982", "Crowd88", "Crusader69Yolo", "Czernysz", "D0n4ldTrump", "DANBHOY67", "DASDIX", "DEYAFALA", "DGonzatti", "DINGGUANWU", "DJWGonTwitch", "DLSI", "DRAKO124rus1", "Dad1227", "Daddyjude", "Daggs", "Dagobert78", "Dakata557", "DakataKalinov", "DamianWKS", "Danaus12", "DarthSidious", "DaviAraujo", "DeathRaven", "DengiDai", "Dereks1", "DiamondHandHodl", "Dicoo", "Dijo", "Dimarikkk1", "DimitriRodrigues", "Diversant88", "Dmitry20279", "DoNCaShHss", "DonMarcos93", "Dontgo", "Dopey5891", "DoresP", "DougBell", "Dracule123", "Dragon310031", "DreamBig", "DshadowwR", "DubaiH", "DudeOG17", "Dukenokaoi", "Dur3mar", "DwaynePipe", "DxxIxxxx", "Dyll92", "DzheykGrin", "DzirtDoUVX", "Ebracham", "Eddyj16", "Edgar38", "Eileithyia", "Einzelganger418", "ElDoctor", "ElZoroLoco", "Eliiia", "Engapenga1", "Erdmandli63", "Estippppp", "Eule86", "Evan87", "Evgen989", "Eyjafjallaj0kull", "F1shru", "FDonky", "FKBurhons", "FL0KI", "Fab17", "Fadimario55", "Falwheel", "Fanosf", "Fatherjack98", "FeelmyAc3", "Felizzaziz", "FilipTrajkoski", "FinnishFish88", "Fish2023v", "Fisherman3105", "Fishnepblba", "Fit4Win1", "Flaeskesteg", "Flavioteixeira91", "Flemari17", "Flopedit88", "Flua67", "FlushTLT1", "FoldMaster05", "FooxSarov", "ForTheirShipping", "Fornaks236", "Fortuneriver", "Fraimworker", "Franbag", "GALAKTIKOS07", "GBL9", "GBpoker", "GDiddz", "GELLERTHOOD365", "GELO000", "GEPETTO1968", "GOATtt", "GOODinBED", "GUHTOH", "GamblerQueen", "Ganyagun", "Gasman1981", "Gilgameshreborn", "GingerAsses", "Gisele92o", "GiveAbrams", "Gjermund", "Glassy", "Gmiapo", "GoingUp", "Gomeslucas28", "GonFrieks", "GoodbyeMoonmen", "GorgonZ0la", "GorillaStyles", "GorinMaxim1", "Gory", "GrandaddyG", "Grantenbee", "GreenZeleboba", "GregPhenomene", "Gror", "Grosuliar1", "GstvPK", "GutsGuGuts", "Gyrandosik", "H0rnyRacc00n", "HADZIBAMBOOBET", "HAULUK171", "HIGHGROVE1", "HOTAR30", "HaSbiKK48772", "Hafez1967", "HairyStone1", "Hamii77", "HappyEnd1945", "HappyHippo6", "Hatypov77", "HazardDS", "HazzeStar", "Health86", "HeateN", "HermosodiA", "HerrSchmidt707", "Hoalan71", "Holdontoyourhat", "Holidayer", "HomeHead", "Honda808", "Howtojump", "Hromko", "Hulipappa", "IIITVMM", "ILUOBSS", "IMABLUFFER72o", "IMMINENCE1", "IMmOrTalBot", "INCREASEXME", "ISLEY", "IVANKL365", "IamCrownupguy", "Icecon", "Ichthyander", "IgorMoraes", "Imas1Marku", "Imso0oh1gh", "InHiatus", "Inaccess", "IncomeNL", "Inkin88", "Inkluz", "Insounich", "Ipochuike", "Itachi23", "Itaferreira", "ItsNotMeNow", "Itsoke", "JAKUBZEHALUK", "JAMA1CA", "JC99", "JIaku762", "JJOkocha", "JJVer777", "JKSol", "JOBATIC", "JOHNTRAIK", "JOSEDANMARK", "JTJoker", "JTomes", "JahSoldier1", "Jahajaa", "JamalHardoev", "JamesSauce", "JammyFker1", "Jangolino", "Janne1962", "Jaxenenko", "Jayg007", "Jedai1", "Jedenactka", "Jeegan", "JensenB42", "Jesus2030return", "JewishWin", "JimiLinny", "Jirissas", "Joecca", "Jokerfish", "Jolle1", "Jonnythefishy", "JoooBoooJooo", "Josilanio365", "JovemGaroto", "JovemIdoso", "Jr2021", "Juancok7424", "Jugoboy777", "Junior2206", "JustMonsters", "JustWokeUp", "KC63", "KEPURO", "KING6979", "KINGMUFFER2", "KORKUT34", "KOSMINASH2", "KOYLIO13", "KPASAB4uK", "Kadiefendi18", "Kahuna", "KaiBT", "KaizerSosa16", "Kannoide", "Karanliginkalbi", "KarpukhinS", "Kaycheto", "KazPokerbro", "Kendo637", "Khwor", "Kikszi", "Killua93", "Kingslayerj", "Kippermuf", "KliX", "Klybinka", "KobaVaan", "KobetsAleksey", "Koersvast", "Kojju", "Kolka9300", "Koneskiz1nn", "Krasnodar84", "KrissCap", "Kucheto888", "Kur40", "KuramaUzumaki", "Kusinator", "LBeloti", "LILA333666", "LLK84", "LLlycTPblu", "LOO5EDEUCE", "LRamilF1", "LUISFERDNAND", "LVCT", "Lakerok12", "LeBoozeux", "LeFisho1", "LeGouns", "Legacies", "LegendCat", "Leggerlista", "Leke89", "LesikMK11", "Lewis13Logan04", "Lilywhite88", "Lintyman", "Loljoe", "London17s", "LordRS96", "LoveSuco", "LoveYr4ips", "Luckybbekker", "Lukas7891", "LukeLamby", "Lydmat111", "MAKEPORN", "MARCFORDE11", "MASTERSH1FU1", "MAVS14", "MEGNYEREM1", "MESKO2009", "MODQPR", "MOGA871", "MSM2oo2", "MUSHKOEZE", "Madds11111", "Magicpat32", "Mahoni88", "MainSweetWinner", "Makla71", "MaksCall", "Malasorte1025", "MansaMussa", "ManuelCh23", "MappaPondiga", "Maritopro8", "Marrogenes", "Martensdebaas", "Mater73", "MaxMnogo", "MaxReppington", "Maxter2001", "McDiv", "McLay77777", "McPokerus", "Meanus", "Mediatedten23", "MelynMefus", "MessiLionel10Arg", "Miasdad2204", "Michele2303", "Mida66", "Mikiru23", "Mil0rd", "Milko", "MillianoSunday", "Mins365", "MissatDrag", "Monkeyspewer", "MontyDArthuR", "Moscowflyer11", "Mougie", "Mount161RUS", "Mozambik87", "MrBabylon23", "MrBingoBongo", "MrCool2016", "MrGold1989", "MrRobot69", "MrSamGold", "Mrdonkthestack", "Mrmouhcine", "MuayTaiPoker", "Mugiwara97", "Mukka64", "Mulkkujussi", "Muris87", "Mydjaxet1", "Mypassword123", "NAKNOATAYE", "NMRF05", "NSIdemo", "NTMATEI1", "Namtur", "NanaPk7", "NanuNanu", "Necula89", "NeilMcauley", "Nerdzzzzz", "Neuropunk1", "Nevkan161", "NicedayxDD", "NickoBelick", "Nicolay1", "NightDreamer", "NikDarZ1", "NikoM", "Ninett90", "NoOneCares", "Notoraye36", "NsFLIPsgfl", "Ntinos13", "Nuncapercoo", "OGTT1", "ONLY1EGG", "ORLOFF", "Oanascumpa82", "Ochia", "Odis", "Oegi", "OggsAces", "Ogito82", "Oilyrags", "OkAndyOk", "Oktaij", "Olidzsii", "Olivebet", "Onaroll", "Onbekend", "OneDayFly", "OniNeSmogli", "OnkelAdolf888", "Osse06", "Oumuamua", "Ovod1972", "OziTom", "PACOBG", "PATtheBOSSyes", "PGFKILL", "PMIpro", "POKERmonPE", "PUBG18kHOURS", "Pannekoekie", "Panzerfranz1", "PapiCSK", "Parisian1985", "PascalR1", "PayToPlay", "PblboJloB21", "Pchelovod2", "Pecas1977", "PedroLGrande", "Pfffff", "Phantom392", "PharaohsGold", "Philipoker35", "PinkGuy44", "Pioupiou1285", "Player5154496118", "Player7959472736", "Plugandplay", "Pokelare90", "PokerKennyboy", "Pokerfacee10", "Pokerswed64", "Poldinio007", "Popishlepi", "PotDiva", "Pretooxnxx", "ProfitFlow", "Prosharker", "Purger999", "Py6Jlbb", "QDECHuK", "QED19681", "Qstik", "Quantumin30sec", "R0BoAfuLL", "R0m1k", "RANCHER8", "RAZZ966", "REAL1PLAYER", "REPRESENT", "RICHI33", "RIVET393", "RKA2", "RONII10", "ROSENBOOM", "RaiseInFace", "RaiseorCheck", "Ramses23000", "RamzeSPL1", "Ramzes71711", "Ranko77", "Raudmoldheia", "Rayancrescenzo", "Rdownes12", "ReVoPoKerPRo", "RealSeriouslee", "RealityMaker", "RedBourbon", "RedLeone", "ReggaeShark420", "Regoslav", "Rekoon", "Remo500", "RetrNSK", "Rex666", "Ricky37", "Rickypower", "Rigodad", "Rimicu", "RiverD", "Road2Diva", "Rodorias", "RomPro", "RomeRRo223", "RonnieFerrari", "Rottweil01", "RoyalAces28", "Royalxflushx6", "Ruby2003", "RumbleFish01", "RunTheGame01", "Runasmmbue", "Ruopolo82", "Ryancash71", "SAMA1973", "SCRGE", "SDSnew777", "SHARKTEETH1", "SIDCITY", "SINYA73", "SLAVCHEV1", "SMOK3Y911", "SSS1ba", "SSTRIKKERR1", "STAM1908", "STEFAN7207", "STILINSKI", "SZILARD333", "Sabadell22", "Saimon777", "Saka7", "SalemRus", "Salesman", "Samafa78", "Samurai3Jack", "SarperBABA", "SasatoMyth", "Sasheoka", "Saywhatusee123", "Schweigen", "SciAthlete", "ScoobyDoo277", "Sen10", "Sentinel", "Serafinas", "Serega13899", "SergeySied1", "SeviS", "Shad0w87", "Shaft1807", "She3py", "Shikz", "ShortyXL1", "Shtet1hajdut", "Sikveland", "Sillpamcho", "SimonBarnes", "Sinclair4", "SirMixStar", "SittingOut666", "Skaffa", "SkillGrinder", "Slaamdunkk", "SlanSan", "Slimblim", "Slimjim1967", "Smilestile", "SnappyGator", "Soagshka", "SochiLove", "Sosnin1", "SouDonCesao", "SouGermann1", "Soundblast1", "Sowa31", "Sparrwo", "Speakme8", "Spiridon64", "Splitpot", "Spurned1", "Square7", "StavS", "StazioneSud", "Stellof", "Stereohaus", "Steven116", "StevenBrunson", "Stoyanyu", "SunWillShine", "Sunny07", "SuperZigmund", "Supernova44", "Sutt66", "Svandyk", "Svetka2020", "SwissUndertaker", "SxWiik", "SystemOFaClown", "Szhovits", "Szitu16", "T1Kdnk", "TESerg1", "THEIOSVERIKOS", "TLK661", "TOB1C4O1", "TOPyKMaKTO", "TY3CEMb", "TaHaTocuk", "Tadukas77777", "Tafgai057", "Tanajurex", "Tashchoo", "Tatar94", "TeacherMW", "TeamAB", "TechnisiaN", "Tedonetime", "Th3K1ng0fQu33n5", "ThEnik2", "TheBlackFish1", "TheBox", "TheChr0nic", "TheLamboFund", "TheVikingFTW", "Thebestday98", "TheisCook1", "ThisISRobbery", "TingLu", "Tombraider13", "Tomdabomb001", "Tonton3000", "Tonyver", "Tooty", "Trabuada", "Treetop18", "Triamid", "TrufanowA", "TruthAldiNNo1", "TryAgainTmrw", "Tyrlan", "UAdrian", "UberBuR", "UgoDon", "Ulrich1313", "V3NOMx3", "VEVI71", "VINNARE666", "Vaibobo", "Vatyas", "Vegas0105", "Velyov365", "Venata93", "VenenoNando13", "Vgar3333", "Vidok112", "ViktorT9", "Vincenity201", "VinnieRenzo", "Virtuoz928", "VitorDoPoste", "WarlockMage", "Wedr", "Westonpro", "WhySoSerious5", "Wildcad33", "Wildman32", "WinPROfit1", "WinakOt", "WkdWizard", "Woitto", "WolfEE90", "Woodlin61", "Wulsaaan", "XCrixusX", "XMASTREE", "XXxcarlaxXX1", "Xexemeco", "XomesiShark93", "Yahtcman", "Yakovv82", "YannFarrer", "YouAreDeadGG", "YourWhiskeyEyes", "ZaNosov", "Zakerone", "ZeBob2", "ZehFlopinho", "Zekrom1", "Zerocola111", "ZhnaBmzha", "ZimaXoloda", "Zitronell", "Zkander1411", "ZoeHellmuth", "ZoomaR", "Zorronina", "Zotti1983", "Zuf2017", "aceable82", "acefurious63", "acegreat223", "adnanhuseyin", "adrinocrommm", "ael91", "afrodave33", "agentbds", "akkar", "akob24", "aledelia86", "aleksmosin", "alexoldscool3", "alfonsoard", "aliboembahe", "allminenow", "alucardrich", "amobus", "angellast828", "angelsweet79", "annieace1968", "arie1968", "arm0307", "artbrut", "atoo", "auraben", "aurasoma79", "aurora3135", "avagod", "aveclaudenum45", "aveme", "avgcs16", "avk93", "awildjoe", "ayrtonsenna", "azrou1", "b0bP0k3r", "b3droxy", "bablokomneka", "badgerbig711", "badgerhard96", "bakimercimek", "bakolino1988", "baronblue87", "baronfalse69", "battlewithin", "bazaclaire", "behappy777", "bermo62", "bf356110", "bigbossaiva", "bikaparaszt12", "bing66", "bladebig85", "bladebrave143", "bladefalse14", "bladefourth12", "bladehappy575", "bluffingftw", "bobby1210", "bookkooq", "booob", "bossblue24", "bossfinal53", "bossfrank149", "bossgreat331", "botecrs", "bottomsup", "bouncingbadger", "boutzi", "brandten91", "brokericy837", "brokerthird539", "brunecat", "brunoandrease8", "brybry", "bsa1334", "bubka89", "bugfirst174", "bullen123", "bupe87", "button", "buttonround447", "cLaudiu90", "caarmeloo", "cachacorush", "cardgiant78", "carey26", "cartonnpierre", "casajo1971", "cash0utpls", "catcold336", "catmighty59", "catwise133", "cavali9x", "cazamaka2", "celle2tony", "cherryman1", "chico4848", "chikotito", "chrs108", "chv891", "clito200", "clownrapid373", "clownsoft578", "clubfalse16", "clubstable33", "clubwild938", "coalman1981", "colonelsober38", "copybook55", "creamer2", "csorbapeter", "cuboyhalifax", "curwari", "dacebell", "dannylangy", "davirobinho", "davos1974", "dazzyb14", "dblx", "dboneva", "ddaazzjj", "de0m", "demonbluffy699", "demonpoker359", "demonstable437", "derekdougan24", "designblue8", "designbold442", "designloose18", "designsweet756", "devilpurple66", "diamondbig659", "diamondfrank11", "diamondfrank98", "dicebold33", "dicehappy48", "dicewitty836", "diksan1", "discohemuli", "djoke1955", "dmdl33d", "doctorfifth49", "dogbites369", "dogis333", "dougvilla", "dragonglad82", "dragontough674", "dsult", "ducksstaff", "dukefalse972", "dukehungry617", "dukewitty223", "duobaobao", "duple1701", "dyfhghdtrgfhysdf", "dykyiwild", "eaglefast224", "eaglehappy538", "eaglesmart25", "eaglesoft221", "eazyedu", "elbarrio219", "ericos11", "esoen86", "etc3tera", "f0xsarov", "fedriko1972", "fedushenka", "fergak56565", "fightermature72", "fightertrue76", "fightingjay", "fingaz09", "fishbright87", "fjlnn", "fleuflop", "flushfair258", "flushlast569", "focus480", "frankukk123", "freddie10", "frozo94", "fsospartan", "funtone", "galyna54", "gamblerwild383", "gambleryellow73", "gamermature337", "gamertight46", "geesam", "generalbold646", "generalrapid55", "generaltight46", "geotom47", "gerov1t", "ggbbglhfs", "ghnoklpoklmb", "gikalim4", "goddesswaifu", "gogomygame", "goonfishkent", "gorkemmm", "gotovSI", "grantdapa1", "gremio55", "gretter99", "grifter31", "grimeboss", "grinderu", "gulugulu", "gunfirst237", "gunfunny954", "gungreedy797", "gunwitty46", "gustav1959", "harrybluf", "havenuts72o", "hazelaar54", "heartdouble577", "heartfirst53", "hearthonest874", "heartround56", "henceg1970", "hof1234", "hornetthird26", "houseable893", "housescary927", "hustled", "iSmokeMyWeed", "iam1nf", "ianrangers", "igorL1979", "igorpoke", "ilianAA", "ilovepickles", "ilya89nt", "ineedbetari", "insaneheights", "irziya", "islandboynoplay", "jacktriggered46", "jaimy7", "jalenmogbg", "jalmar", "jam1690", "jamieg55", "janemiao", "janika12345", "jarata79", "javascma77", "jayscraic", "jedifast982", "jeszczeKoreq", "jimte", "jjdoe", "joeroe", "john0607", "johnnyturn", "johtajam", "jokercruel536", "jokerjolly49", "jokernor", "jokerround189", "jokersilly5", "jolanda1966", "jonesy8990", "jumbomclooney74", "jumentodoPS", "jumpergreedy31", "jumperloose14", "jupiterbloody13", "jupiterpink484", "justme13", "jutsike", "juvituu", "k1nGd08L", "kadwalk86", "kakadudu007", "kaorle", "karenx83x", "karlos777", "katesman12", "kaufmi", "keff21", "kelly0", "kengus54", "kermantabas", "kevraf", "kidpokerak1", "kingsilver75", "kingtrue25", "kinkysara", "kisu", "kiyababy", "klojo70", "kloppklopp", "knifelazy1766", "knifepurple77", "knifered189", "knightsmart3170", "kokoronashi", "koloskylo86", "komis7", "kospkr", "kostyan3871", "kpkp189", "kraistdelux", "krish212", "kritsakos", "kunedd838", "kuti1092", "kvmvz", "kynder7", "larssonseven", "lavder", "lee1961B210EU", "leksa021", "lemonbloody659", "lemonfirst796", "letitgo44", "letunovskiy1", "lexabij", "lightbluffy14", "lightbright868", "lightfair1224", "lightgreedy412", "lightgrey485", "lightsober698", "lilamlunsoi", "lilkushi7", "lionrapid694", "lionsober85", "loengenheir0", "lombardo33", "louis8cypher", "lovruxa367", "luba313", "lublupelmeni", "luckergoodnn1", "luckyfarquar", "luckygirlwoop", "ludofrik", "luuk1223", "m4gt", "mVizz", "macwariorB", "maddog74", "madvane11", "mainbord1", "maksfaza", "malikecoffee1", "malkena", "mama93", "mando049", "manokwari", "maratalb", "marsbar90", "marshallsilly18", "marvis09", "masterlazy819", "masumoto", "matrixbig91", "matrixpoker10", "maukasta33", "mavun322", "mcman321", "minigu", "mironych781", "mjkoz", "moinogi", "mokumXXXV", "moldexs", "moneytrees12", "monkeyfree816", "monsterwild375", "moonfourth52", "moongutsy88", "moonmighty5", "moonorange686", "moonteal298", "mourashow", "mouseblank732", "mousebronze886", "moush2", "mrbiggut", "mrjeiger", "mrn1c", "mugivaranoluffy", "muhk3l1", "musicBoss1", "mvrmvr2", "naejyy", "namefitz", "nasalo1", "naumandrei", "nefertiti03", "nevadajay51", "newroad1", "ninjabig420", "ninjagiant233", "ninjagreat18", "ninjahot535", "nirgaT", "nkpapu26", "noimo1907", "northoltranger", "nosnurbelyod", "notmycupoftea", "nunugomish1", "oaksey01", "oblofass", "odie", "oli4", "ololosh322", "ondra79", "onetimetel", "orelmaggot1", "orelmaggottt", "orhun", "ossiossporn", "ozanarez", "p0k3rm0n", "paddycom", "panosasd13", "papahappinhooo", "papai89", "papi0624", "paroblods", "paterpieter", "paulrat3", "pavel1697", "pavkom", "perkes31", "peterkoyotes", "pezzmon", "pgaz51", "pigeonloyal93", "pigeonquiet214", "pigeonsilly42", "pilotcruel686", "pilothappy456", "pilottense854", "pinheirodelima", "piotrexik", "pirat372", "piratefalse72", "piratewise186", "pistolfunny166", "pkkk", "playboy34", "player20232", "playergreedy66", "playerplayer5", "playmaker1717", "poehpoeh2", "pokerknight85", "pokkerihai42", "polkanow", "ponyfinal741", "ponyfrank135", "popeyeyid", "ppppppp9999", "prACEnuts", "professorwild23", "progutsy6139", "proteal561", "prowise6711", "pst0", "psychess", "psychiadelic123", "ptrenka", "punterspicy34", "putkari215", "qq2bb", "quadsgreedy287", "quadshonest354", "quadsteal36", "quadstense32", "queenhot8478", "rafaelm123", "ragnor8921", "rainbowbold7", "rainbowfast823", "rainbowhard957", "ramos16", "raplunks33", "rattyface", "ravendark843", "ravenwise95", "reaper88", "renegat35", "rense", "reswobKK", "rexgutsy2", "rexgutsy993", "ribas6", "ricardo0207", "rim5511", "riverratje57", "rizhiy28a", "robotmad353", "robotsleepy945", "robotwitty1", "rookie365", "roopop", "rosebean", "roseblue933", "rosemad1427", "rosemighty945", "ru11ogin", "rumpybumpy", "runnerfunny66", "ryan134556", "rz78", "s1ll1e", "sadieadler", "sahsa89", "sailorcool93", "sailorquiet61", "sal69", "saltamonte17", "samuraifair261", "sances991", "sany7788", "sashaviktor1", "saturnfifth45", "saturnicy583", "sbmccoy", "sbtyra", "scariik", "scarponi", "sedu", "sembach12", "seppuku11", "septemberholding", "serask01", "serpentiro", "setsoft78", "seumadrugakiing", "shaddowgti", "shankie123456", "sharklovely245", "sharkpoker878", "sheriffbold656", "shipfluffy935", "shipgreat567", "shipgutsy119", "shipgutsy99", "shiplazy754", "shiployal64", "shipround732", "sicjoker1", "sittler2727", "sjaak171", "sjakkmats", "skywhy1", "slanting99", "sluglips", "smokesomerolled", "smokytoon", "sms231", "snooopyyhooo", "soberano01", "solarbeam", "sontrips", "sormland1", "spadecold231", "spadedropped98", "spadeloyal628", "spadescary568", "spadeteal89", "specis13", "spicedark46", "spooncool966", "spoonred784", "spoonroyal462", "spoontriple412", "ssakisso1", "ssn1968x", "sspokerman", "stakerdark356", "stakerlast244", "stammers007", "steviebyg", "straightmad71", "streetglad612", "streetsilly651", "strongimam", "stupi011", "su8459", "sukkel313", "sullyhey", "sunaggro42", "sunblue921", "sundmc21", "sungreedy317", "sunloyal139", "sunplatinum22", "sunwich20021", "superakkayht", "supermind", "supreeme", "svettpeis", "swanjoe", "swissmeister", "sxprst12", "syleiman19901", "synn007", "szipi76", "t1ma", "tableguy101", "takis50s", "tankgirl53", "tbone893", "tchjka", "teachermad91", "teranjevolova", "terje61", "tern10", "teteuc22", "texasvalisi", "teytey27", "themaniacblair", "thenastychef", "thesmokeybandit", "thissiane", "tigergreen715", "tigerpurple759", "time52", "timus87", "tinkertom", "titanmature27", "tjallex", "tlsv", "toad36", "tommo6343", "tonyjmlfc", "tosuave", "trapper1353", "trevorrr007", "tripshappy79", "tripsready855", "tryyourlook", "tso1337", "tvoypapa", "twixxx999", "ugur0101", "unionjack2014", "universecool47", "upc47", "ursinus89", "uvikb5555", "vaa17", "vadyaryadom", "vangdamus", "vanko0707", "vengrusss", "vidmiz25", "viking7119", "viole", "visioneery1", "vitbog73", "vizaut2", "vladimirLVA", "vladinhoA", "vnelly", "vnh1234", "w3ak3n3r1", "wagwanfam", "wallsmart78", "wardenround916", "wardentable34", "weng5000", "whaleable28", "whalelele", "whalepoker32", "whelleshepard", "whitesugar2", "wifesdeboss2", "wisard58", "wolftomee", "wowureallysuck", "wx3000", "xCJIECAPbx", "xTrevPay67x", "xVektoRx", "xXCrownUpGuyXx", "xmandmytro2", "xqzzz", "xxGame", "xxcampbell44", "xxxmohxxx", "yohohomofo", "yrik78871", "yurakulik", "yurovskikh1", "zebby1", "zero525", "zetem000", "zhoponuh", "ziglizagli", "zikaferal", "znakrazmuta", "zoiandre", "zora1964", "zumutz"}
var agentsList []string
var proxyList []string

func init() {
	var err error
	rand.Seed(time.Now().Unix())
	agentsList, err = readLines("user_agents.txt")
	if err != nil {
		log.Fatal(err)
	}
	proxyList, err = readLines("proxies.txt")
	if err != nil {
		log.Fatal(err)
	}
}

func Get(player, room string) (string, error) {
	var result string
	var err error
	for count := 0; count < 10; count++ {

		result, err = requestStats(player, room)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}
	return result, err
}

func requestStats(player, room string) (string, error) {
	defer timeTrack(time.Now(), "requestStats")
	tmpUrl := fmt.Sprintf("https://sharkscope.com/poker-statistics/networks/%s/players/%s?&Currency=USD", room, player)

	proxyUrl, _ := url.Parse(proxy())

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	req, err := http.NewRequest("GET", tmpUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("User-Agent", userAgent())
	req.Header.Add("Accept-Language", "en")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Username", "")
	req.Header.Add("Password", "")
	req.Header.Add("X-Requested-With", "XMLHttpRequest")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Referer", "https://sharkscope.com/")

	response, err := httpClient.Do(req)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	return string(body), nil
}

func userAgent() string {
	return agentsList[rand.Intn(len(agentsList))]
}

func proxy() string {
	var p = "http://"
	if len(proxyList) == 1 {
		p += proxyList[0]
	} else {
		p += proxyList[rand.Intn(len(proxyList))]
	}
	return p
}

func RandPlayer() string {
	return playersList[rand.Intn(len(playersList))]
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func timeTrack(start time.Time, name string) {
	fmt.Printf("%s заняло %s\n", name, time.Since(start))
}

func GetInfoIP() (string, error) {
	for count := 0; count < 10; count++ {
		result, err := requestInfo()
		if err != nil {
			fmt.Println(err)
			continue
		}

		reg, _ := regexp.Compile("{.*}")
		result = reg.FindString(result)
		//println(result)

		return result, err
	}
	return "", errors.New("не удалось получить ответ по истечении 10 попыток")
}

func requestInfo() (string, error) {
	tmpUrl := "https://coding.tools/my-ip-address"

	proxyUrl, err := url.Parse(proxy())
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	httpClient := &http.Client{
		Timeout: 10 * time.Second,
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	response, err := httpClient.Post(tmpUrl, "application/json", nil) //bytes.NewBuffer([]byte("queryIp=''"))
	if err != nil {
		return "", err
	}

	//response, err := httpClient.Do(req)
	//if err != nil {
	//	return "", err
	//}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	return string(body), nil
}
