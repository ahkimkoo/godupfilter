package main

import (
	"github.com/ahkimkoo/godupfilter/shingle"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

var (
	numThreads = runtime.NumCPU()
	task       = make(chan map[uint32]bool, numThreads)
	done       = make(chan bool, numThreads)
	numRuns    = numThreads
	shex       map[uint32]bool
	totalTimes = 100000
)

func init() {
	sx := []byte(`随着城市化进展加速，拆迁成为了我们身边一种非常常见的现象。  但你也许不了解的是，拆迁这桩生意背后，几乎每一个环节都隐藏着暴利因子，有些拆迁项目利润率有的能达到300%-400%，拆迁公司操作一个项目短短几月就能赚几百万……如此暴利的产业，究竟是怎样做到的？  ■口述 /高祥 某木材厂老板，拆迁公司股东访谈-整理 /巴丽  字数：2867字 阅读时间：约4分钟  ▽▽▽  我的老家在陕西咸阳乾县马连镇，是个远近闻名的“拆迁专业镇”，我们镇的人基本都在从事跟拆迁相关的产业，比如我，除了有一个自己的木料场在经营外，也会接一些拆迁工程自己做，甚至入股别人的拆迁工程。  我们镇不大，但仅木材厂就有差不多300多家，是附近有名的旧木料市场，其中很多木料都是从西安、咸阳的拆迁工地上拉回来的。当地很多木材厂的老板也都跟我一样，至少有两个身份：木材经销商和拆迁公司股东，而且因为这两个生意挣了大钱的不在少数。  随着这两年西安咸阳市大规模的拆迁工程越来越多，挣到钱的不仅仅只有拆迁公司，从工程发包方、中间人、承包工程的拆迁公司、实际做工程的公司，到挖机出租商、渣土运输车、旧材料市场，都成为拆迁利益链条中获利的一环。  毫不夸张的说，拆迁生意就是一座暴利的“金矿”。  1  “拆迁利益链”  算起来我做拆迁工程也有十多年了，其中干的最多的就是城中村拆迁。    请点击此处输入图片描述  拆迁招标分为两种情况，首先都是由有资质的拆迁公司承包拆迁工程。  接下来，会根据不同工程具体状况，分两种方式操作：一种是工程利润非常大的，工程发包方会按每平米几十元的收费标准进行招标，各家拆迁公司会根据评估结果决定是否竞标（当然，最后中标的肯定都是有背景的关系户）；对于一些利润不大但拆迁又比较费劲的项目，工程发包方会按每平方米几十元的费用标准给中标的拆迁公司进行补贴。  无论以上哪种情况，按照规定，拆迁公司须向发包方交纳每平方米15-20元的保证金。说白了其实就是押金，一是防止拆迁公司中途撂挑子不干，二是出了安全事故时用来赔偿。  拆迁公司拿到工程后，便租赁挖机，组织农民工、联系渣土车主。拆迁公司收益主要来自拆下来的钢筋、门窗等，这笔费用加上发包方给的费用（或刨去上交给发包方的费用）及挖机出租方、农民工及渣土车主相应费用后，就是拆迁公司的纯利。这是个稳赚不赔的生意。只要不出安全事故，肯定都能赚钱。  拆下来的东西中，钢筋和门窗最值钱。  一般的城中村拆迁，每平方米可产生20公斤左右的钢筋，拆下的门框可卖5-6元/平方米。以一个10万平方米左右的城中村为例，承包价格为15元/平方米，拆后每平方米值26-27元，刨掉承包费、人工费、拉土费、机械费用后，一个城中村拆下来，短短几个月，轻轻松松就能赚到好几百万。  2  “怎么揽活？”  拿工程，主要看关系。有的工程先由专门负责拆迁的国有性质的经营公司拿到，然后再转包给拆迁公司。有门路的则可以直接从相关部门接活、找发包公司，否则，就得靠中间人介绍了。  当然，好处费是少不了的，而且一般这个费用还都不低。具体数额因人而异、因具体工程而异。关系越硬，回扣就越少，利润就越大。  拿到活后，拆迁公司就开始评估。评估就是定价，一平方米多少钱，主要是看建筑里边货多货少。“货”主要是指钢筋、门窗，还有砖头、电表、铁、地板等等。  按照规定，拆迁公司必须具有行业资质。那些没有资质的拆迁公司，为了揽到活，必须找一家有资质的公司挂靠，实在找不到挂靠单位的，往往会冒险借别人的执照。有一回承包华阴一个活，我们一次拿了16个照，都是借别人的。  还有一部分工程，会政府相关部门交给开发商，再由开发商发包给拆迁公司。想拿这样的工程，你必须和开发商很熟，或者是能搭上比较可靠的关系，不然很拿到手。  3  “怎么找人合伙？”  评估完成，确定价格后，拆迁公司开始寻找合伙人，一般都是由大股东或者牵头拿到工程的人去找股东入伙。  根据工程量和押金量，股东可多可少。一般10万算一股（也就是想入股最少得10万），多的可以几百万。2008年时，我和别人在西安北郊承包了一个活，30万平方米，需交押金500万元。那次我们找了18个股东。  利润也是按出资比例分配，大股东分的更多，因为工程是大股东揽过来的。有时候他说要给介绍人提成，可能好几百万，直接从“空中”提走，这些钱都不入账。实际上这些钱一部分给了介绍人，一大半很可能就是大股东自己直接拿走了。  动迁时，拆迁公司的股东既能挣固定工资，又能挣提成。  一般情况下，固定工资每月大概是1800元，迁走1户还能额外挣最高上万的提成。也有些不摊成本不出人，只拿钱不干活的干股东，这些人一般是介绍人、村干部甚至还有黑社会。不要这些股东，很可能你这活都干不成。  在我们马连镇，入股拆迁的最少也有1000人，其中资金实力较为雄厚、可以牵头或当大股东的有一两百人，小股东则有八九百人。一次能拿出二三百万的起码有200人。  在马连镇，一天攒个一两千万不是问题。只要有活干，股东多的是，马上就能拿钱入股。镇上的信合和邮政储蓄一天内钱就能给提空了，经常能看到县上的运钞车一天来回几趟地往镇上运钱。  4  “拆迁环节如何挣钱？”  找齐股东，交了押金，就可以开始进入拆迁环节，包括动迁、安置村民。产生的所有费用基本都由开发商承担，政府只是出面协调和村民的关系。拆迁改造前开发商会先把钱给政府，政府再补偿给拆迁户。    请点击此处输入图片描述  这部分费用（包括回迁分配的住房、门面房等）是相对固定的。建成后的楼房先安置村民。所以如果楼盖得低了，开发商利润也就薄了，这也是为什么开发商都喜欢盖高楼的原因，因为楼越高利润越大。  大部分村民在拿到补偿、安置款后，都能顺利搬走。  当然，也有少数钉子户。为了赶工期，只能强制拆迁。这期间产生的一切费用，比如把屋里的东西损坏了，先由一些政府部门出面协调、谈判，最后实在不行就由开发商拿钱赔偿。  拆迁暴利行业的利润到底有多大？可以说，一般都是投多少，挣多少，活多的话挣得更多。  我算了一笔细账：2009年拆迁西安南郊某村，共50万平方米，发包方给出的价格是30元/平方米，中间人拿走20元/平方米后，到他们手里只剩下10元，就这样也挣钱。  其中，挖机每平方米3.5元，共175万元；清理垃圾按每平方米10元付给渣土车主，共500万元；卖旧钢筋、门窗、椽子等共收入750万元，再刨除管理费10万元，这个工程总共赚了565万元，工期只用了不到3个月。  这个工程我们交了600万元的押金，其中最大的股东出了300万元，其余有拿100万元的，也有出10万、20万的。我当时只入股了10万元，我们这样的小股东既在工地干活，拿工资，还能参与分红。这一个工程做下来，也赚了将近10万元。  3个月，将近100%的利润！而且一般的拆迁工程，工期都不会超过3个月，真真是一门时间短、见效快、利润高的好生意。  干拆迁，利润低于20%就算赔钱。但100%的利润也不算高，最高的300%、400%都有。  在拆迁公司的眼里，工地上遍地是宝，包括电子表，三相电表，黑电表，地板砖，废铁皮，废塑料管，铝线等废料处理，就是一笔大买卖。保守来算，一个工程仅块就有几百万的收入。  拆迁既然这么暴利，许多人当然也趋之若鹜，纷纷欲挤进这个一夜暴富的行当。  好多公司为了揽到活，最后都不惜动刀子，工地上经常能看见两帮人拿着棍棒、刀子对打，说到底，都是看中了这个行当的暴利。  我总结，胆大、心黑、有钱是干这一行必备的素质。但随着这几年房地产行业不景气，各地城市化进入尾声，拆迁工程数量也没有前几年那么多了，估计再过两年想进入的人也会少很多。  变革家生意笔记：  1、 拆迁招标分为两种情况：发包方给钱或收费，但无论哪种情况都有钱赚，而且基本都是有背景的关系户拿到工程；  2、 拿工程，主要看关系。有门路的可以直接找发包公司，否则，就得靠中间人介绍，而且好处费不低。总的说来，关系越硬，回扣就越少，利润就越大。  3、 利润按出资比例分配，大股东分的更多，因为工程是大股东揽过来的。也有些只拿钱不干活的干股东，这些人一般是介绍人、村干部甚至还有黑社会。  4、 干拆迁，利润低于20%就算赔钱，但100%的利润也不算高，最高的300%、400%都有。 一项工程，仅废料处理就至少有几百万的收入。  最近随着变革家给大家带来的商机内幕越来越多，后台有很多人表示希望能分享一下他们的业内经验。变革家在这里正式向各位业内人发出邀请，热烈欢迎大家和变革家几十万中产以上投资用户分享您的干货，也许这些人就是您的未来合作伙伴。`)
	shex = shingle.Shingling(sx, 3, 0)
}

func worker(no int) {
	for shnw := range task {
		// log.Printf("no.%d get task", no)
		for i := 0; i < totalTimes/numRuns; i++ {
			shingle.Similarity(shex, shnw)
		}
	}
	done <- true
}

func main() {
	if len(os.Args) > 2 {
		tta, err1 := strconv.Atoi(os.Args[1])
		if err1 == nil && tta > 0 {
			totalTimes = tta
			log.Printf("total times: %v\n", totalTimes)
		}

		tnum, err := strconv.Atoi(os.Args[2])
		if err == nil && tnum > 0 {
			numThreads = tnum
			log.Printf("numThreads: %v\n", numThreads)
		}
	}

	runtime.GOMAXPROCS(numThreads)

	s1 := []byte(`央视网消息：2017年2月23日至24日，中共中央总书记、国家主席、中央军委主席习近平到北京市考察城市规划建设。他强调，北京城市规划建设和北京冬奥会筹办工作是当前和今后一个时期北京市的两项重要任务，抓好城市规划建设，着眼精彩非凡卓越筹办好北京冬奥会，努力开创首都发展更加美好的明天。  　　做好城市规划，确定中国的城市未来走向，习近平有着系统化的思路，并在不同场合有过很多的阐述。  　　合理布局 规划先行  　　城市规划在城市发展中起着战略引领和刚性控制的重要作用，做好规划，是任何一个城市发展的首要任务。  　　2月24日下午，习近平在人民大会堂北京厅主持召开北京城市规划建设和北京冬奥会筹办工作座谈会。他指出，城市规划在城市发展中起着重要引领作用。北京城市规划要深入思考“建设一个什么样的首都，怎样建设首都”这个问题。  　　事实上，规划先行的理念一直贯穿于习近平关于城市建设的思路中。2014年2月，习近平在北京考察，首站选择在北京市规划展览馆。他指出，考察一个城市首先看规划，规划科学是最大的效益，“规划失误是最大的浪费，规划折腾是最大的忌讳。”  　　在此次北京城市规划建设和北京冬奥会筹办工作座谈会上，习近平要求，把握好战略定位、空间格局、要素配置，坚持城乡统筹，落实“多规合一”形成一本规划、一张蓝图，着力提升首都核心功能，做到服务保障能力同城市战略定位相适应，人口资源环境同城市战略定位相协调，城市布局同城市战略定位相一致，不断朝着建设国际一流的和谐宜居之都的目标前进。  　　建设北京城市副中心，是党中央一项重要决策。2月24日上午，通州城东，副中心行政办公区现场指挥部，习近平察看规划沙盘，观看视频短片，了解副中心建设理念、目标定位、文化保护等情况。  　　“站在当前这个时间节点建设北京城市副中心，要有21世纪的眼光。”习近平指出，规划、建设、管理都要坚持高起点、高标准、高水平，落实世界眼光、国际标准、中国特色、高点定位的要求。  　　“不但要搞好总体规划，还要加强主要功能区块、主要景观、主要建筑物的设计，体现城市精神、展现城市特色、提升城市魅力。”  　　执行规划 一以贯之  　　好的蓝图更需要好的执行。近几年，随着我国城镇化的急剧发展，“千城一面”“一任领导一任规划”“政绩工程”等尴尬问题不断凸显，如何更好地、更有力地执行城市规划成为一大难题，对此，习近平认为保证规划的严肃性和权威性非常重要，而实现这一目标的重要途径之一就是立法。  　　2013年12月，改革开放以来首次举行的中央城镇化工作会议中也更加明确了城市规划要保持连续性，“不能政府一换届、规划就换届。编制空间规划和城市规划要多听取群众意见、尊重专家意见，形成后要通过立法形式确定下来，使之具有法律权威性。”  　　舟山，被称为“千岛之城”，是我国首个以群岛建制的地级市。2015年5月，习近平在参观舟山城市展示馆时就指出：“舟山群岛新区规划要确保法定效力。”  　　要发挥城市规划的法定效力，健全而行之有效的决策和执行机制必不可少。2015年12月，中央城市工作会议在北京举行。会议上，习近平总书记强调，要全面贯彻依法治国方针，依法规划、建设、治理城市，促进城市治理体系和治理能力现代化。要健全依法决策的体制机制，把公众参与、专家论证、风险评估等确定为城市重大决策的法定程序。要深入推进城市管理和执法体制改革，确保严格规范公正文明执法。  　　2016年2月中共中央、国务院印发《关于进一步加强城市规划建设管理工作的若干意见》。四大着力点、七大重点任务、三十项具体意见，问诊把脉当下的城市规划管理。《意见》中特地强调，经依法批准的城市规划，是城市建设和管理的依据，必须严格执行。  　　此次，在北京城市规划建设和北京冬奥会筹办工作座谈会上习近平再次强调，总体规划经法定程序批准后就具有法定效力，要坚决维护规划的严肃性和权威性。  　　不忘初心 人民城市为人民  　　北京市规划展览馆位于北京的前门东大街，里面陈列着几块《北京城市总体规划》展板。  　　2014年2月，在这里考察的习近平看到展板上绿色越来越多后，说，“网上有人给我建议，应多给城市留点‘没用的地方’，我想就是应多留点绿地和空间给老百姓。”  　　城市规划建设做得好不好，最终要用人民群众满意度来衡量。  　　如何提高人民群众的满意度，在此次的座谈会上，习近平提出，要以北京市民最关心的问题为导向，以解决人口过多、交通拥堵、房价高涨、大气污染等问题为突破口，提出解决问题的综合方略，要健全制度、完善政策，不断提高民生保障和公共服务供给水平，增强人民群众获得感。  　　“要坚持人民城市为人民。”习近平多次表明要下大力气根治“城市病”增强民众获得感的决心——  　　要加大大气污染治理力度，应对雾霾污染、改善空气质量的首要任务是控制PM2.5；  　　要把解决交通拥堵问题放在城市发展的重要位置，加快形成安全、便捷、高效、绿色、经济的综合交通体系；  　　解决群众住房问题是一项长期任务……我们必须下更大决心、花更大气力解决好住房发展中存在的各种问题。  　　城镇化进程中的核心是人，“城镇化不是土地城镇化，而是人口城镇化”“户籍人口城镇化率直接反映城镇化的健康程度”。对此，习近平强调：“总的政策要求是全面放开建制镇和小城市落户限制，有序放开中等城市落户限制，合理确定大城市落户条件，严格控制特大城市人口规模，促进有能力在城镇稳定就业和生活的常住人口有序实现市民化，稳步推进城镇基本公共服务常住人口全覆盖。”  　　按照“十三五”规划，2020年要实现常住人口城镇化率达到60%左右，户籍人口城镇化率达到45%左右，这就意味着，我国每年至少要有近两千万人口从农村进入城镇工作生活，而做好城市规划，依法落实城市建设，转变城市发展方式，提高城市发展质量和人民满意度，则是推动社会经济健康发展，提升人民群众获得感的必然选择。`)
	s2 := []byte(`随着城市化进展加速，拆迁成为了我们身边一种非常常见的现象。  但你也许不了解的是，拆迁这桩生意背后，几乎每一个环节都隐藏着暴利因子，有些拆迁项目利润率有的能达到300%-400%，拆迁公司操作一个项目短短几月就能赚几百万……如此暴利的产业，究竟是怎样做到的？  ■口述 /高祥 某木材厂老板，拆迁公司股东访谈-整理 /巴丽  字数：2867字 阅读时间：约4分钟  ▽▽▽  我的老家在陕西咸阳乾县马连镇，是个远近闻名的“拆迁专业镇”，我们镇的人基本都在从事跟拆迁相关的产业，比如我，除了有一个自己的木料场在经营外，也会接一些拆迁工程自己做，甚至入股别人的拆迁工程。  我们镇不大，但仅木材厂就有差不多300多家，是附近有名的旧木料市场，其中很多木料都是从西安、咸阳的拆迁工地上拉回来的。当地很多木材厂的老板也都跟我一样，至少有两个身份：木材经销商和拆迁公司股东，而且因为这两个生意挣了大钱的不在少数。  随着这两年西安咸阳市大规模的拆迁工程越来越多，挣到钱的不仅仅只有拆迁公司，从工程发包方、中间人、承包工程的拆迁公司、实际做工程的公司，到挖机出租商、渣土运输车、旧材料市场，都成为拆迁利益链条中获利的一环。  毫不夸张的说，拆迁生意就是一座暴利的“金矿”。  1  “拆迁利益链”  算起来我做拆迁工程也有十多年了，其中干的最多的就是城中村拆迁。    请点击此处输入图片描述  拆迁招标分为两种情况，首先都是由有资质的拆迁公司承包拆迁工程。  接下来，会根据不同工程具体状况，分两种方式操作：一种是工程利润非常大的，工程发包方会按每平米几十元的收费标准进行招标，各家拆迁公司会根据评估结果决定是否竞标（当然，最后中标的肯定都是有背景的关系户）；对于一些利润不大但拆迁又比较费劲的项目，工程发包方会按每平方米几十元的费用标准给中标的拆迁公司进行补贴。  无论以上哪种情况，按照规定，拆迁公司须向发包方交纳每平方米15-20元的保证金。说白了其实就是押金，一是防止拆迁公司中途撂挑子不干，二是出了安全事故时用来赔偿。  拆迁公司拿到工程后，便租赁挖机，组织农民工、联系渣土车主。拆迁公司收益主要来自拆下来的钢筋、门窗等，这笔费用加上发包方给的费用（或刨去上交给发包方的费用）及挖机出租方、农民工及渣土车主相应费用后，就是拆迁公司的纯利。这是个稳赚不赔的生意。只要不出安全事故，肯定都能赚钱。  拆下来的东西中，钢筋和门窗最值钱。  一般的城中村拆迁，每平方米可产生20公斤左右的钢筋，拆下的门框可卖5-6元/平方米。以一个10万平方米左右的城中村为例，承包价格为15元/平方米，拆后每平方米值26-27元，刨掉承包费、人工费、拉土费、机械费用后，一个城中村拆下来，短短几个月，轻轻松松就能赚到好几百万。  2  “怎么揽活？”  拿工程，主要看关系。有的工程先由专门负责拆迁的国有性质的经营公司拿到，然后再转包给拆迁公司。有门路的则可以直接从相关部门接活、找发包公司，否则，就得靠中间人介绍了。  当然，好处费是少不了的，而且一般这个费用还都不低。具体数额因人而异、因具体工程而异。关系越硬，回扣就越少，利润就越大。  拿到活后，拆迁公司就开始评估。评估就是定价，一平方米多少钱，主要是看建筑里边货多货少。“货”主要是指钢筋、门窗，还有砖头、电表、铁、地板等等。  按照规定，拆迁公司必须具有行业资质。那些没有资质的拆迁公司，为了揽到活，必须找一家有资质的公司挂靠，实在找不到挂靠单位的，往往会冒险借别人的执照。有一回承包华阴一个活，我们一次拿了16个照，都是借别人的。  还有一部分工程，会政府相关部门交给开发商，再由开发商发包给拆迁公司。想拿这样的工程，你必须和开发商很熟，或者是能搭上比较可靠的关系，不然很拿到手。  3  “怎么找人合伙？”  评估完成，确定价格后，拆迁公司开始寻找合伙人，一般都是由大股东或者牵头拿到工程的人去找股东入伙。  根据工程量和押金量，股东可多可少。一般10万算一股（也就是想入股最少得10万），多的可以几百万。2008年时，我和别人在西安北郊承包了一个活，30万平方米，需交押金500万元。那次我们找了18个股东。  利润也是按出资比例分配，大股东分的更多，因为工程是大股东揽过来的。有时候他说要给介绍人提成，可能好几百万，直接从“空中”提走，这些钱都不入账。实际上这些钱一部分给了介绍人，一大半很可能就是大股东自己直接拿走了。  动迁时，拆迁公司的股东既能挣固定工资，又能挣提成。  一般情况下，固定工资每月大概是1800元，迁走1户还能额外挣最高上万的提成。也有些不摊成本不出人，只拿钱不干活的干股东，这些人一般是介绍人、村干部甚至还有黑社会。不要这些股东，很可能你这活都干不成。  在我们马连镇，入股拆迁的最少也有1000人，其中资金实力较为雄厚、可以牵头或当大股东的有一两百人，小股东则有八九百人。一次能拿出二三百万的起码有200人。  在马连镇，一天攒个一两千万不是问题。只要有活干，股东多的是，马上就能拿钱入股。镇上的信合和邮政储蓄一天内钱就能给提空了，经常能看到县上的运钞车一天来回几趟地往镇上运钱。  4  “拆迁环节如何挣钱？”  找齐股东，交了押金，就可以开始进入拆迁环节，包括动迁、安置村民。产生的所有费用基本都由开发商承担，政府只是出面协调和村民的关系。拆迁改造前开发商会先把钱给政府，政府再补偿给拆迁户。    请点击此处输入图片描述  这部分费用（包括回迁分配的住房、门面房等）是相对固定的。建成后的楼房先安置村民。所以如果楼盖得低了，开发商利润也就薄了，这也是为什么开发商都喜欢盖高楼的原因，因为楼越高利润越大。  大部分村民在拿到补偿、安置款后，都能顺利搬走。  当然，也有少数钉子户。为了赶工期，只能强制拆迁。这期间产生的一切费用，比如把屋里的东西损坏了，先由一些政府部门出面协调、谈判，最后实在不行就由开发商拿钱赔偿。  拆迁暴利行业的利润到底有多大？可以说，一般都是投多少，挣多少，活多的话挣得更多。  我算了一笔细账：2009年拆迁西安南郊某村，共50万平方米，发包方给出的价格是30元/平方米，中间人拿走20元/平方米后，到他们手里只剩下10元，就这样也挣钱。  其中，挖机每平方米3.5元，共175万元；清理垃圾按每平方米10元付给渣土车主，共500万元；卖旧钢筋、门窗、椽子等共收入750万元，再刨除管理费10万元，这个工程总共赚了565万元，工期只用了不到3个月。  这个工程我们交了600万元的押金，其中最大的股东出了300万元，其余有拿100万元的，也有出10万、20万的。我当时只入股了10万元，我们这样的小股东既在工地干活，拿工资，还能参与分红。这一个工程做下来，也赚了将近10万元。  3个月，将近100%的利润！而且一般的拆迁工程，工期都不会超过3个月，真真是一门时间短、见效快、利润高的好生意。  干拆迁，利润低于20%就算赔钱。但100%的利润也不算高，最高的300%、400%都有。  在拆迁公司的眼里，工地上遍地是宝，包括电子表，三相电表，黑电表，地板砖，废铁皮，废塑料管，铝线等废料处理，就是一笔大买卖。保守来算，一个工程仅块就有几百万的收入。  拆迁既然这么暴利，许多人当然也趋之若鹜，纷纷欲挤进这个一夜暴富的行当。  好多公司为了揽到活，最后都不惜动刀子，工地上经常能看见两帮人拿着棍棒、刀子对打，说到底，都是看中了这个行当的暴利。  我总结，胆大、心黑、有钱是干这一行必备的素质。但随着这几年房地产行业不景气，各地城市化进入尾声，拆迁工程数量也没有前几年那么多了，估计再过两年想进入的人也会少很多。  变革家生意笔记：  1、 拆迁招标分为两种情况：发包方给钱或收费，但无论哪种情况都有钱赚，而且基本都是有背景的关系户拿到工程；  2、 拿工程，主要看关系。有门路的可以直接找发包公司，否则，就得靠中间人介绍，而且好处费不低。总的说来，关系越硬，回扣就越少，利润就越大。  3、 利润按出资比例分配，大股东分的更多，因为工程是大股东揽过来的。也有些只拿钱不干活的干股东，这些人一般是介绍人、村干部甚至还有黑社会。  4、 干拆迁，利润低于20%就算赔钱，但100%的利润也不算高，最高的300%、400%都有。 一项工程，仅废料处理就至少有几百万的收入。  最近随着变革家给大家带来的商机内幕越来越多，后台有很多人表示希望能分享一下他们的业内经验。变革家在这里正式向各位业内人发出邀请，热烈欢迎大家和变革家几十万中产以上投资用户分享您的干货，也许这些人就是您的未来合作伙伴。`)
	shg1 := shingle.Shingling(s1, 3, 0)
	shg2 := shingle.Shingling(s2, 3, 0)
	distance := shingle.Similarity(shg1, shg2)
	log.Printf("%f\n", distance)

	for i := 0; i < numThreads; i++ {
		go worker(i)
	}

	log.Printf("Start work")

	t0 := time.Now()

	for i := 0; i < numRuns; i++ {
		task <- shg1
	}
	close(task)

	for i := 0; i < numThreads; i++ {
		<-done
	}

	t1 := time.Now()
	log.Printf("finish, cost: %v", t1.Sub(t0))
}
