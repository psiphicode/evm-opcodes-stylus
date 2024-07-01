package main

import "github.com/holiman/uint256"

var num_u256 int = 100
var test_cases_u256 []*uint256.Int = []*uint256.Int{
    &uint256.Int{5881117533349185536, 11670090440229713920, 13203022344362997760, 907046039771983872},
    &uint256.Int{9879106341554468864, 18426313596868921344, 13341837974018277376, 2702473745593540608},
    &uint256.Int{2964094150432147456, 16028721796185765888, 17290273215306110976, 9975033854299715584},
    &uint256.Int{4794998364457859072, 2090535761227569152, 11841457105720793088, 5113779939285731328},
    &uint256.Int{1912634685002143744, 5696061458869659648, 2548263142389760, 14988881153822134272},
    &uint256.Int{17845674914207827968, 3055967396049899520, 2964559355909101568, 8922454298826436608},
    &uint256.Int{114091447764944896, 13050928109526134784, 16485579909423818752, 5784954265745496064},
    &uint256.Int{14040265333504741376, 11924790292924796928, 12978593677781854208, 7527410907786369024},
    &uint256.Int{14060066879833638912, 16274258391281045504, 3573267300580732928, 13975459364057153536},
    &uint256.Int{17083786717491886080, 18291247392172324864, 17218355559793584128, 18180931993166784512},
    &uint256.Int{13002931025797128192, 12598041309804515328, 16588798064595591168, 8018263811336437760},
    &uint256.Int{9123706460447991808, 11860628423936681984, 4654708644181035008, 11454699030145589248},
    &uint256.Int{7197809615521468416, 3806763727664619520, 9463061652286662656, 7497019834524708864},
    &uint256.Int{13108259805512853504, 14889085966651807744, 18309553875850557440, 17618087381409837056},
    &uint256.Int{9568844120065585152, 10914976222214801408, 16025550111113955328, 17023888652673548288},
    &uint256.Int{15136077842639908864, 14991152008735688704, 11463439214542327808, 2686983207831554048},
    &uint256.Int{4503320607211231232, 15788572708460718080, 2963073776344895488, 16304793201260902400},
    &uint256.Int{11732437830316515328, 4770032976626817024, 5565030306910136320, 7001471009920278528},
    &uint256.Int{6840164872286531584, 4574289314651365376, 18244296551354484736, 15891778243218612224},
    &uint256.Int{2483228620370860032, 2077779797044891648, 12379320355633801216, 6768241935266938880},
    &uint256.Int{5091669962607329280, 1379174383957207040, 7366704721136556032, 10389345568388384768},
    &uint256.Int{16312257383486459904, 16426417116220784640, 18149224650398175232, 6706116174416097280},
    &uint256.Int{14285188308136900608, 8247587646071838720, 1121223673824415744, 3583503779472889856},
    &uint256.Int{10567797846370328576, 17428041463486932992, 7213479767364603904, 15748740435133167616},
    &uint256.Int{803107882041606144, 1952655553387268096, 18019411088132257792, 16829940797672820736},
    &uint256.Int{14539544169378715648, 12682298300751763456, 5346094347285657600, 17691640648652673024},
    &uint256.Int{17909556773588832256, 9441791903407429632, 3749079425785147392, 5704761038019477504},
    &uint256.Int{6932537685158694912, 17314851266928535552, 13613827770248200192, 6223243437500493824},
    &uint256.Int{17761233788653582336, 460395208841277440, 3797519310879942656, 1747655337877612544},
    &uint256.Int{12162179535394189312, 1064060885486024704, 16645069158968434688, 13552034508233111552},
    &uint256.Int{2959172983439130624, 18128066600411156480, 2324536509528283136, 1501066429770463232},
    &uint256.Int{9754420969539395584, 17278161826969933824, 8903939841961533440, 9750163961602793472},
    &uint256.Int{10050719422228862976, 2442821242588878848, 7029409890458843136, 8880519731217328128},
    &uint256.Int{12413845099522918400, 9800306238602887168, 1862658340932792320, 444220318219104256},
    &uint256.Int{15940629135956797440, 15529530037660264448, 1702095460608804864, 9516785628657829888},
    &uint256.Int{3080902713012791296, 3711153277742716928, 13089600191250157568, 14511568385785976832},
    &uint256.Int{4335427756454875136, 8036125655626938368, 15437239993199181824, 3943463319614482432},
    &uint256.Int{14089730125521938432, 11542259589089665024, 11205088607585775616, 9218469775375120384},
    &uint256.Int{16657437554106761216, 16331914413991780352, 5983332963771240448, 5114737267802300416},
    &uint256.Int{3711807800714586112, 11700085567969947648, 10462543118215897088, 8484148224900935680},
    &uint256.Int{14379548034710253568, 14956338473570181120, 10865704505679306752, 17868413152302383104},
    &uint256.Int{7424146457202298880, 4424122125563246592, 9643809190221901824, 14246272009269661696},
    &uint256.Int{5050204708176128000, 2654598817576259584, 764801304848617472, 11867375333320491008},
    &uint256.Int{9190918484321546240, 17514341532022175744, 17335609381290946560, 14624557679216265216},
    &uint256.Int{7123603615607042048, 16812874693372516352, 9038710918661894144, 7594970457313292288},
    &uint256.Int{7232626915495428096, 7924281910037985280, 9583745814474876928, 1120985061116469248},
    &uint256.Int{12522704173323286528, 17903413333903245312, 5934665118490038272, 8028548299250771968},
    &uint256.Int{16443702305651245056, 9857074947692730368, 7316380238663012352, 5661518777657233408},
    &uint256.Int{17663046527846961152, 7947068821464371200, 15606922051579009024, 18223883124021786624},
    &uint256.Int{3421949644327663616, 14004787301253181440, 3591344587973154816, 5753921274569572352},
    &uint256.Int{16996032258036269056, 12863117379403001856, 4973716984884641792, 10047020915633956864},
    &uint256.Int{9930163010676948992, 2714588258602387456, 6552583302577453056, 3042413765667592192},
    &uint256.Int{5684780877838931968, 16161894543492904960, 2758871806243168256, 14665547374623899648},
    &uint256.Int{10968235889802702848, 13172980589090342912, 2985284672387395584, 13784922755476332544},
    &uint256.Int{16207464204004855808, 14187744111920236544, 8191277033449482240, 6698965459803052032},
    &uint256.Int{13157977274294503424, 5921830973117159424, 7351276972827348992, 14842199665123604480},
    &uint256.Int{8091358228736944128, 17235047613602865152, 6965725893154056192, 337037944851603456},
    &uint256.Int{16084767161717346304, 7779610167256317952, 14130584771631529984, 9373849983852580864},
    &uint256.Int{14948918512405925888, 14039040889621964800, 17119127255679911936, 12066561102515142656},
    &uint256.Int{12610212361559414784, 3579635428653256704, 2354160322247696384, 12297606455621578752},
    &uint256.Int{15503905279420413952, 8024777553928687616, 15912902423166369792, 9832864080450672640},
    &uint256.Int{680179463124711424, 621124276936767488, 14241283696511543296, 12182367040413114368},
    &uint256.Int{13162375588842270720, 12543428800964685824, 11204280430356256768, 12119170368657174528},
    &uint256.Int{9429345543959064576, 4276244881228533760, 10083110984480065536, 17440738702392750080},
    &uint256.Int{2178505446695274496, 9531081856514416640, 10428723610497474560, 13770477747724617728},
    &uint256.Int{14064133325412747264, 2969609666933817344, 9966390447568969728, 541074841474267136},
    &uint256.Int{12006020915786405888, 18255164795355930624, 4554020215373383680, 5444956815946899456},
    &uint256.Int{2988081910739884032, 15020785770889805824, 5039642013472620544, 1073155617298837504},
    &uint256.Int{6433810984881420288, 6098668128643614720, 7149458072151404544, 15473711874206582784},
    &uint256.Int{12706743584959115264, 17575392740813668352, 17750015982158981120, 5362137695493681152},
    &uint256.Int{3575440494180638720, 10723873736987033600, 368918926599467008, 10548426525273141248},
    &uint256.Int{6910896089745152000, 5776101941239941120, 7315381370266636288, 10748432657199489024},
    &uint256.Int{1001175449600182272, 14426894228805732352, 14281822967930302464, 2890289498328836096},
    &uint256.Int{15016901917558429696, 5521013834027634688, 17681757214271934464, 11217836335715336192},
    &uint256.Int{15063989119728918528, 5219789484076210176, 9761872560335161344, 10916726248002037760},
    &uint256.Int{1176592955739344896, 17406068051620642816, 278231106100092928, 11765533419607506944},
    &uint256.Int{8308406877193996288, 856804983789131776, 4748727485386160128, 15975898998197972992},
    &uint256.Int{2601225982649679872, 2642914982733035520, 5331656806759536640, 6279025171153090560},
    &uint256.Int{8198505696515667968, 2575647404289499136, 5249242453519785984, 1052423710261587968},
    &uint256.Int{11659526738938066944, 15483398007528157184, 4343036620376832000, 4772112876594178048},
    &uint256.Int{15814476094625349632, 12316757844352669696, 6738588530163691520, 11933274754201329664},
    &uint256.Int{15918035924498262016, 17905541716111730688, 13264583613250156544, 2993981271295723520},
    &uint256.Int{717064401028878336, 14595259376214315008, 2556300704396744704, 2650040603168129024},
    &uint256.Int{3818474322630703104, 18309997161896374272, 12722233319198562304, 5219166199591655424},
    &uint256.Int{2407702078230263808, 13864518445248352256, 1572864555432161280, 12940938268741433344},
    &uint256.Int{2672350157403635712, 18183700537609764864, 17751059296046032896, 17528846217626277888},
    &uint256.Int{14080774170942203904, 12494581924542812160, 8333842480419596288, 5960760038160922624},
    &uint256.Int{18191720217538988032, 1555893177015181312, 7438697412601266176, 11518154311878242304},
    &uint256.Int{1934973455268069376, 7801593540211851264, 6245430379045761024, 6027442257476648960},
    &uint256.Int{14649446115805362176, 8602291539980425216, 3075583410053595136, 3862890513204811776},
    &uint256.Int{15880140130077339648, 13511651789889292288, 17963293284360714240, 14575502851307823104},
    &uint256.Int{5224353467446749184, 5167335417082724352, 3455747722235557888, 16469549416450254848},
    &uint256.Int{3344193078320410624, 3022737606468683776, 3147662937533536256, 1816603145049929728},
    &uint256.Int{2811083221334628352, 2295365001017628672, 3456160716578699264, 14879272830503575552},
    &uint256.Int{1105405530589061120, 4554749473403289600, 17793994539409471488, 2031132054272931840},
    &uint256.Int{15811504691666534400, 2846025507415076864, 17481208657331109888, 17446076851131545600},
}

var num_u64 int = 100
var test_cases_u64 []uint64 = []uint64{
    581937585264236544,
    9855087055923689472,
    6875387778842877952,
    8436526837985587200,
    9843688811271245824,
    14450249530569082880,
    10377580084245606400,
    8178338813435156480,
    1493474139041456128,
    5746466027431411712,
    16019017979066564608,
    3583389946872954880,
    16583697499751843840,
    5767245692807559168,
    10113976910765398016,
    14994245377285083136,
    14485945308644212736,
    17967580348862455808,
    5373623266974652416,
    2928550489106489344,
    1974585440508499968,
    8565895168199829504,
    12777197602412853248,
    9304908666713786368,
    10177682396630429696,
    16365008954358626304,
    10104376533807423488,
    13126103619395950592,
    2715665090750550016,
    396757141170610176,
    1375893829041950720,
    8187671918131451904,
    6785431385520398336,
    12030723834414516224,
    17922630377381167104,
    3871402545711933440,
    622095103673526272,
    2163299054511798272,
    17052560124854159360,
    3170940628167008256,
    11423292476476631040,
    12022755118797496320,
    2999590103198289920,
    7291833532133986304,
    6599232086943666176,
    13970253875384997888,
    4759441076446910464,
    4095556410968483840,
    9910442891921459200,
    4286153478838599680,
    2003139211806994432,
    8887453945524955136,
    77517315930370048,
    10362906820345671680,
    14504900169651855360,
    314556144655618048,
    16495981754711912448,
    16021220900283260928,
    10457947606524618752,
    17644497430969843712,
    15793095601155403776,
    17060976001271441408,
    3479704816755621888,
    15123135560801732608,
    16702105287323852800,
    7698441075944925184,
    16234135535619506176,
    7694094412311326720,
    8331620179114964992,
    4123904125439481856,
    4873232005438803968,
    5496350137467865088,
    15333265582865139712,
    7562209264883152896,
    13764651651520303104,
    28030428466900992,
    1444198064302981120,
    15893552719344748544,
    10465269021973721088,
    13248101983470469120,
    9510653290676713472,
    357390122931163136,
    12382255689863501824,
    9983422680243447808,
    12452543019220416512,
    17740408481378531328,
    497062071831328768,
    16988012134249910272,
    3657111127719759872,
    7403376704811130880,
    990483873230956544,
    5239550706952947712,
    12662401956689739776,
    2003808211088322560,
    459936503062421504,
    18306624108360007680,
    18067642673886939136,
    16339571865478928384,
    13460380181038231552,
    4877441424554641408,
}
