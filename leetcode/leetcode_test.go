package leetcode

import (
	"fmt"
	"testing"
)

func Test1046(t *testing.T) {
	S := "0110"
	N := 10
	fmt.Println(queryString(S, N))
}

func TestBin(t *testing.T) {
	//for i:=0; i<=5;i++ {
		fmt.Println(bin(5))
	//}
}

func Test748(t *testing.T) {
	liscensePlate := "1s3 PSt"
	words := []string{"step","steps","stripe","stepple"}

	fmt.Println(shortestCompletingWord(liscensePlate, words))

}

func Test976(t *testing.T) {
	fmt.Println(largestPerimeter([]int{3,6,2,3}))
}

func Test1371(t *testing.T) {
	fmt.Println(findTheLongestSubstring("eleetminicoworoep"))
}

func Test1218(t *testing.T) {
	arr := []int{1,5,7,8,5,3,4,2,1}
	dif :=   -2
	fmt.Println(longestSubsequence(arr, dif))
}

func Test402(t *testing.T) {
	num := "99641436378815361153471302158193420182863684789411484994976484827114595334610042544056442370530816060833617030976813134098793056155103202008549344446519354408307307071055065112738442020228471569394796174150323080161225901964338837341524253243218509500254619223683091799365677720582389568156585225666197123093377871100002481402486219837255411382162499321193416524972275273471969155848742457476556433737281147710021781210134765321761285612276511917324552585569882156635094670362653567596144728653795007023230933817566104488637696450166087905100823699425798763598444326069357052842379918535855296915760054459317433521878778171811081076593166663090948029793113626852462712388116483774713426183911481230884393594249331828165503798269634244430773693033882708000249632850148799859322024693146577635543114657662418998860517525989192973250701631765598465053097616804817344343895016724561947860836117504915797011185132674255278236597746042138768473723059825948301565719437610732907662545499042953499866813741157301003371005200992314265077531029437948931255617153417148822355928318598517533241719641002712204874161001604269216566928220767474713135516717997491363360204764154264989004671363541097433484822118483642107547658581450616821769964767032521138851570822729134762460014265433227201724724004338494552397280090568164786109721571436206198382814849033856987338787473335772666933218810822482848994610491705665155516384799459418594559136827941106387689501641851101743298582575466303864906673788496628288920867422193950180810131396612913851112593807649152972068279299934113463669714575613645929365652921808836725682390026075559320995704880149764583379697505303474550029059828116836469203370428449330442281563135568935742669243344218603994417955703485059862132359688776290378210392955310785874528205203788559715493852405991380290274268143557970398441851157977520689440430265144029789788511042795879174567381358510694749512938934687979305099149575464220629804942550564164786808856897809863824121659548034395539735407069279457678613909222371848892294754933299091164656871086269084324529512544747434123547189729993758337622038098699448815701644934651292719067683227727438808955969543542319197883567369733867364250353136697865107182282929655918362211832327827571354787535611501731943856155003853732339819594939524719169561110698571676562329360803282215467534058504728127731515598941143637827010955579092451405821352126706550438315176049692316210490899702613078702535716735901806171522853021035597316703390478571485677998207922773938829371460838611214446417528913575284776737837046439695408523434414916342979688820197836458637694991540998291690345194205452439239827382953039810367712244590155940394387554911786652478111954297185544106384174592451680875083737874735810068767866214924634885513828808880161930987276602570872860752119813042414550396358433893592777541756673206882876746731707766966268096104320061937913505893028833592137540396064375155513979764728180927083060481127522118240026140625647313783901073938419240249929000962722034273952683635919540169732220854978101308126446671885186032295490845060116567165945677975672981321362161949418405852378788584602802612398876874288293756055559457538271197205867506313677160755990736347314715042607243878693780144368083800080967842966193539823770427967091132770230485036143223363387876244958899577538069175123004651952588711287008791159682042581943812962882375293348462523257081140457567348612069746943329842264291823570671268374580651696311114624358304235261945894627668267192756606441264485628097480920062857007640396910214970556623416565940789636657349735150043836194242061994234044262604284350296258397208287158735477739515615890093167555389262170576609082365199242352356197706754361085079177223144662701424848070607319078068303190442737202186364818021792860690571733432439513976759807778513151206801184300729685910765785586373831699595178352610150383283823456881293647763022411686252640648120690251120902631370825525354213297549430441989419362406888242180413640397005462289002837178086683143441254722528075315187910994986929463063282350677644105312484770818851268755183086729904524488901102287310169865855725358976453628171038414004415469635124255044890245890050115901243603489384920067923087045070616429510114587493955384903357111302068595548921504222171096098548413208088831560744996899783844118318185694142620796984004522106434428513215881883542758888862576036415421097762413907290417004936441609238204617100586876487061497586106631983740139555573272626681186969272113315348553052708453716313010811194726904231406455432865684477036960953564406390115786323388585604716504384778912812410729908949581143722120318954849846535676912868526526501078193502393524062471534154104899815734648650035608611113327222040864146091286020205304970098510045582130989981665393076480660907742469107193219475618455618115516353495211289597815564506193368287178714208989206470099207227171770619580227427772058576958549342547850566371060314330889132466260972915500785842700966615103949831075688522846389635990078358138687466663099265099431775674237640711466272609872329090894406587154198409486434056948991642623725868520261081714501891452704954562834244485695899485150794033902595303371632597184940525684558272222395813587950566598836575728711404672894869851301199508345442816914540274231773573695049117433232750564343477296571911336451338765122801905492189124021699698020217831160061375249740348841211772476455089061870953510480256335713228323198782026742817220321247980121667780800877801219532811542139900480803615083739957513418528009253849655053312995534574307148952727627870318872325094411860749809155407484065987101730385346571248798467212335910821152286411077915790397497756477613051365987943518909759211252763081626026136209474490841118337332773116122063152414208776801671614382203998310801791046109980464795153775904284579208046765170299376571712696359391195309011046580945099118345329164807866461624513459858969478261348365746242842254100449074846018162381649508771205692387943049083877156128753239386498305599949138477358461424273464036997642435352074743094695564535693378173888280633866732018710701060752702258884562187458492514181027419045608607139753797741693225900923436163273291784047946102859573341135995351940672974945745062320931107916232460722010886651827074516009065280667168017782964663521168472263155891094369134584611694802620433621767214124173962636180142978128638945692419270222518432363382128100260544917455244318162619360808797214154001396840051520865249909119773623276044783996235484958441702533661095335337458603732924068113476544273220040621287278168707393471504842692312354782265568742305367773635557008065688109790648713350572351799924638273829816187626279342407486758617884199669669286080608957640162096427744397522103026413782698158732581790000716751490076906346484023835702438474105176931779065689980130347837155056303467742499515965713045957954225592059807462917282749105358673064716135765849677591608061323905019687616579401117839719269327243007586938365568212311638431283680946079388989080798521721770825311237382299640977231722390040018733060008726711369177955792504805871660952275133036361448257222162174106121886956846208577175900217031085260775753651365765038925717954695019720235653672968689019573262654460436772900765775615489257834882352941349073672575670561593061387879337673233294306479935031268311515186416299622966578517978675818927585118344348361158710756313053131716293124192982037977789379782122120656399498488608931743952536041546453299501041577456229618221253519224906611827751220393777623642577532653929191439603183004880021982807536023221789599010502125687724004685177438516674638976736887749480118357141229355178588718777866510629202733751110559334924038607709059709853979249569510212755627954315025008066453716096825677236680969921750877126730256949811077056975031686370565845816981036167892330455103497165407984322792515265566483796338273488042877728447328933645773410093062365682687268013318931065552717013674172822704288279197461978805944285413220284999303849740540429893025407810120053701999064303195562726870079068213843151094378846458471168159763363401468459072474435300433314015701363633705309153196187013664717617975618648227816754951474354742056233896619815305871556180590934191775446450232435064334173434855333465262160341517250209548644211312373841441024747539900101488865742679168673356769004244781832745045012713439497231232255815861738982590755401780194874615548229070120796893835181030047378827641086164272219294123942746140207443292075817414598536256892540490923602419336928186124051416665048479530882042184097629985897052425322145715174649893481917612568426372077919256931921063600255204010662044398922537796993713110889134889921360833579323314386803074533058134342770923839546994120322442157750203621967931319597649960815556196358566683782572730174920215034531104191490057838260392829741446722127017532444082857280503217574522928285094747407153894570747792487061998260753833304433675066923630595212677695003060727653119915126939127827754432456052655283764591328484359469704894122366077507922825301623961196207923544095047285011474898262448957681893278273601046641810135121516552187096005252171171905022763076761687166299014789581539855448453229411352775826042558462563147630238335355859149814380543807473386539264830261256996173935860136236427622918234260408201158550118527706241993700526213016072648406003487895118011337828945314863348154387066988573131543747121745028364818130265528614742576976975564213718421245904443000581698214695522541683926528961160986876871840844632069685227319014872180179370554032205521013345746425253133686231659075343389374580200717637698542920298315739628019867736462368334051114029380922339886663078026309916370486909128253195100898377068612057592121356555290537815049586626181680384845905180029133497372417653664436161971980137048236053329456957495141918670077299206755740534997886723627476115663811233372206043170460623060506091246306386543951687123557178508806912199010111871"
	k := 1000
	fmt.Println("result", removeKdigits(num, k))
}

func TestClear(t *testing.T) {
	a := "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010016042692165669282207674747131355167179974913633602047641542649890046713635410974334848221184836421075476585814506168217699647670325211388515708227291347624600142654332272017247240043384945523972800905681647861097215714362061983828148490338569873387874733357726669332188108224828489946104917056651555163847994594185945591368279411063876895016418511017432985825754663038649066737884966282889208674221939501808101313966129138511125938076491529720682792999341134636697145756136459293656529218088367256823900260755593209957048801497645833796975053034745500290598281168364692033704284493304422815631355689357426692433442186039944179557034850598621323596887762903782103929553107858745282052037885597154938524059913802902742681435579703984418511579775206894404302651440297897885110427958791745673813585106947495129389346879793050991495754642206298049425505641647868088568978098638241216595480343955397354070692794576786139092223718488922947549332990911646568710862690843245295125447474341235471897299937583376220380986994488157016449346512927190676832277274388089559695435423191978835673697338673642503531366978651071822829296559183622118323278275713547875356115017319438561550038537323398195949395247191695611106985716765623293608032822154675340585047281277315155989411436378270109555790924514058213521267065504383151760496923162104908997026130787025357167359018061715228530210355973167033904785714856779982079227739388293714608386112144464175289135752847767378370464396954085234344149163429796888201978364586376949915409982916903451942054524392398273829530398103677122445901559403943875549117866524781119542971855441063841745924516808750837378747358100687678662149246348855138288088801619309872766025708728607521198130424145503963584338935927775417566732068828767467317077669662680961043200619379135058930288335921375403960643751555139797647281809270830604811275221182400261406256473137839010739384192402499290009627220342739526836359195401697322208549781013081264466718851860322954908450601165671659456779756729813213621619494184058523787885846028026123988768742882937560555594575382711972058675063136771607559907363473147150426072438786937801443680838000809678429661935398237704279670911327702304850361432233633878762449588995775380691751230046519525887112870087911596820425819438129628823752933484625232570811404575673486120697469433298422642918235706712683745806516963111146243583042352619458946276682671927566064412644856280974809200628570076403969102149705566234165659407896366573497351500438361942420619942340442626042843502962583972082871587354777395156158900931675553892621705766090823651992423523561977067543610850791772231446627014248480706073190780683031904427372021863648180217928606905717334324395139767598077785131512068011843007296859107657855863738316995951783526101503832838234568812936477630224116862526406481206902511209026313708255253542132975494304419894193624068882421804136403970054622890028371780866831434412547225280753151879109949869294630632823506776441053124847708188512687551830867299045244889011022873101698658557253589764536281710384140044154696351242550448902458900501159012436034893849200679230870450706164295101145874939553849033571113020685955489215042221710960985484132080888315607449968997838441183181856941426207969840045221064344285132158818835427588888625760364154210977624139072904170049364416092382046171005868764870614975861066319837401395555732726266811869692721133153485530527084537163130108111947269042314064554328656844770369609535644063901157863233885856047165043847789128124107299089495811437221203189548498465356769128685265265010781935023935240624715341541048998157346486500356086111133272220408641460912860202053049700985100455821309899816653930764806609077424691071932194756184556181155163534952112895978155645061933682871787142089892064700992072271717706195802274277720585769585493425478505663710603143308891324662609729155007858427009666151039498310756885228463896359900783581386874666630992650994317756742376407114662726098723290908944065871541984094864340569489916426237258685202610817145018914527049545628342444856958994851507940339025953033716325971849405256845582722223958135879505665988365757287114046728948698513011995083454428169145402742317735736950491174332327505643434772965719113364513387651228019054921891240216996980202178311600613752497403488412117724764550890618709535104802563357132283231987820267428172203212479801216677808008778012195328115421399004808036150837399575134185280092538496550533129955345743071489527276278703188723250944118607498091554074840659871017303853465712487984672123359108211522864110779157903974977564776130513659879435189097592112527630816260261362094744908411183373327731161220631524142087768016716143822039983108017910461099804647951537759042845792080467651702993765717126963593911953090110465809450991183453291648078664616245134598589694782613483657462428422541004490748460181623816495087712056923879430490838771561287532393864983055999491384773584614242734640369976424353520747430946955645356933781738882806338667320187107010607527022588845621874584925141810274190456086071397537977416932259009234361632732917840479461028595733411359953519406729749457450623209311079162324607220108866518270745160090652806671680177829646635211684722631558910943691345846116948026204336217672141241739626361801429781286389456924192702225184323633821281002605449174552443181626193608087972141540013968400515208652499091197736232760447839962354849584417025336610953353374586037329240681134765442732200406212872781687073934715048426923123547822655687423053677736355570080656881097906487133505723517999246382738298161876262793424074867586178841996696692860806089576401620964277443975221030264137826981587325817900007167514900769063464840238357024384741051769317790656899801303478371550563034677424995159657130459579542255920598074629172827491053586730647161357658496775916080613239050196876165794011178397192693272430075869383655682123116384312836809460793889890807985217217708253112373822996409772317223900400187330600087267113691779557925048058716609522751330363614482572221621741061218869568462085771759002170310852607757536513657650389257179546950197202356536729686890195732626544604367729007657756154892578348823529413490736725756705615930613878793376732332943064799350312683115151864162996229665785179786758189275851183443483611587107563130531317162931241929820379777893797821221206563994984886089317439525360415464532995010415774562296182212535192249066118277512203937776236425775326539291914396031830048800219828075360232217895990105021256877240046851774385166746389767368877494801183571412293551785887187778665106292027337511105593349240386077090597098539792495695102127556279543150250080664537160968256772366809699217508771267302569498110770569750316863705658458169810361678923304551034971654079843227925152655664837963382734880428777284473289336457734100930623656826872680133189310655527170136741728227042882791974619788059442854132202849993038497405404298930254078101200537019990643031955627268700790682138431510943788464584711681597633634014684590724744353004333140157013636337053091531961870136647176179756186482278167549514743547420562338966198153058715561805909341917754464502324350643341734348553334652621603415172502095486442113123738414410247475399001014888657426791686733567690042447818327450450127134394972312322558158617389825907554017801948746155482290701207968938351810300473788276410861642722192941239427461402074432920758174145985362568925404909236024193369281861240514166650484795308820421840976299858970524253221457151746498934819176125684263720779192569319210636002552040106620443989225377969937131108891348899213608335793233143868030745330581343427709238395469941203224421577502036219679313195976499608155561963585666837825727301749202150345311041914900578382603928297414467221270175324440828572805032175745229282850947474071538945707477924870619982607538333044336750669236305952126776950030607276531199151269391278277544324560526552837645913284843594697048941223660775079228253016239611962079235440950472850114748982624489576818932782736010466418101351215165521870960052521711719050227630767616871662990147895815398554484532294113527758260425584625631476302383353558591498143805438074733865392648302612569961739358601362364276229182342604082011585501185277062419937005262130160726484060034878951180113378289453148633481543870669885731315437471217450283648181302655286147425769769755642137184212459044430005816982146955225416839265289611609868768718408446320696852273190148721801793705540322055210133457464252531336862316590753433893745802007176376985429202983157396280198677364623683340511140293809223398866630780263099163704869091282531951008983770686120575921213565552905378150495866261816803848459051800291334973724176536644361619719801370482360533294569574951419186700772992067557405349978867236274761156638112333722060431704606230605060912463063865439516871235571785088069121990101118710"
	fmt.Println(clear(a))
}

func Test874(t *testing.T) {
	coms := []int{1,2,-2,5,-1,-2,-1,8,3,-1,9,4,-2,3,2,4,3,9,2,-1,-1,-2,1,3,-2,4,1,4,-1,1,9,-1,-2,5,-1,5,5,-2,6,6,7,7,2,8,9,-1,7,4,6,9,9,9,-1,5,1,3,3,-1,5,9,7,4,8,-1,-2,1,3,2,9,3,-1,-2,8,8,7,5,-2,6,8,4,6,2,7,2,-1,7,-2,3,3,2,-2,6,9,8,1,-2,-1,1,4,7}
	obs := [][]int{{-57,-58},{-72,91},{-55,35},{-20,29},{51,70},{-61,88},{-62,99},{52,17},{-75,-32},{91,-22},{54,33},{-45,-59},{47,-48},{53,-98},{-91,83},{81,12},{-34,-90},{-79,-82},{-15,-86},{-24,66},{-35,35},{3,31},{87,93},{2,-19},{87,-93},{24,-10},{84,-53},{86,87},{-88,-18},{-51,89},{96,66},{-77,-94},{-39,-1},{89,51},{-23,-72},{27,24},{53,-80},{52,-33},{32,4},{78,-55},{-25,18},{-23,47},{79,-5},{-23,-22},{14,-25},{-11,69},{63,36},{35,-99},{-24,82},{-29,-98},{-50,-70},{72,95},{80,80},{-68,-40},{65,70},{-92,78},{-45,-63},{1,34},{81,50},{14,91},{-77,-54},{13,-88},{24,37},{-12,59},{-48,-62},{57,-22},{-8,85},{48,71},{12,1},{-20,36},{-32,-14},{39,46},{-41,75},{13,-23},{98,10},{-88,64},{50,37},{-95,-32},{46,-91},{10,79},{-11,43},{-94,98},{79,42},{51,71},{4,-30},{2,74},{4,10},{61,98},{57,98},{46,43},{-16,72},{53,-69},{54,-96},{22,0},{-7,92},{-69,80},{68,-73},{-24,-92},{-21,82},{32,-1},{-6,16},{15,-29},{70,-66},{-85,80},{50,-3},{6,13},{-30,-98},{-30,59},{-67,40},{17,72},{79,82},{89,-100},{2,79},{-95,-46},{17,68},{-46,81},{-5,-57},{7,58},{-42,68},{19,-95},{-17,-76},{81,-86},{79,78},{-82,-67},{6,0},{35,-16},{98,83},{-81,100},{-11,46},{-21,-38},{-30,-41},{86,18},{-68,6},{80,75},{-96,-44},{-19,66},{21,84},{-56,-64},{39,-15},{0,45},{-81,-54},{-66,-93},{-4,2},{-42,-67},{-15,-33},{1,-32},{-74,-24},{7,18},{-62,84},{19,61},{39,79},{60,-98},{-76,45},{58,-98},{33,26},{-74,-95},{22,30},{-68,-62},{-59,4},{-62,35},{-78,80},{-82,54},{-42,81},{56,-15},{32,-19},{34,93},{57,-100},{-1,-87},{68,-26},{18,86},{-55,-19},{-68,-99},{-9,47},{24,94},{92,97},{5,67},{97,-71},{63,-57},{-52,-14},{-86,-78},{-17,92},{-61,-83},{-84,-10},{20,13},{-68,-47},{7,28},{66,89},{-41,-17},{-14,-46},{-72,-91},{4,52},{-17,-59},{-85,-46},{-94,-23},{-48,-3},{-64,-37},{2,26},{76,88},{-8,-46},{-19,-68}}
	fmt.Println(robotSim(coms, obs))
}

func Test_16_18(t *testing.T) {
	pattern := "a"
	value := "zqvamqvuuvvazv"
	fmt.Println(patternMatching(pattern, value))
}