package znr_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	znr "github.com/billglover/zn-reader"
)

func TestKnownPhrases(t *testing.T) {
	vl := znr.VocabList{
		znr.Vocab{Writing: "你"},
		znr.Vocab{Writing: "是"},
		znr.Vocab{Writing: "好"},
		znr.Vocab{Writing: "友"},
		znr.Vocab{Writing: "你好"},
	}

	tr := znr.NewTrie()
	for _, v := range vl {
		tr.Insert(v.Writing)
	}

	cases := []struct {
		txt string
		tr  znr.Trie
		out []string
	}{
		{txt: "你好世界", tr: tr, out: []string{"你好"}},
		{txt: "你好世界，你今天怎么样？", tr: tr, out: []string{"你好", "你"}},
		{txt: "你是谁？", tr: tr, out: []string{"你", "是"}},
		{txt: "我是你的好朋友", tr: tr, out: []string{"是", "你", "好", "友"}},
	}

	for _, c := range cases {
		known, err := tr.KnownPhrases(c.txt)
		if err != nil {
			t.Error(err)
		}

		if reflect.DeepEqual(known, c.out) == false {
			t.Errorf("%v != %v", strings.Join(known, ","), c.out)
		}
	}
}

func BenchmarkKnownPhrases(b *testing.B) {
	vl := znr.VocabList{
		znr.Vocab{Writing: "你"},
		znr.Vocab{Writing: "是"},
		znr.Vocab{Writing: "好"},
		znr.Vocab{Writing: "友"},
		znr.Vocab{Writing: "你好"},
	}

	tr := znr.NewTrie()
	for _, v := range vl {
		tr.Insert(v.Writing)
	}

	for t := 8; t < 12; t++ {
		l := 1 << t
		b.Run(fmt.Sprintf("%03d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = tr.KnownPhrases(txt[0:l])
			}
		})
	}
}

func TestTrieInsert(t *testing.T) {
	cases := []struct {
		txt string
		f   bool
	}{
		{txt: "北京", f: true},
		{txt: "中国", f: true},
		{txt: "北", f: true},
		{txt: "安静", f: false},
		{txt: "", f: false},
	}

	tr := znr.NewTrie()

	for _, c := range cases {
		if c.f == false {
			continue
		}
		tr.Insert(c.txt)
	}

	for _, c := range cases {
		if f := tr.Find(c.txt); f != c.f {
			t.Errorf("%s found: %t expected %t", c.txt, f, c.f)
		}
	}
}

// Source: https://zh.wikipedia.org/wiki/中华人民共和国
var txt = `中华人民共和国，简称中国或大陸[註 1]，是一个位於欧亚大陆东部的社会主义国家及主权国家，法定首都为北京[15]。中国領土陸地面積估約
960萬平方公里，是世界上纯陸地[註 14]面積第二大、陸地[註 15]面積第三大、總面積第三大或第四大的國家[註 16][16]，其分為23個省份[註 17]、
5個自治區、4個直轄市和2個特別行政區。中國地势西高东低而呈現三级阶梯分布，大部分地区属于溫帶、副熱帶季风气候，地理景致與氣候型態丰富多樣，
有冰川、丹霞、黃土、沙漠、喀斯特等多种地貌[17]，中国北方分布有乾草原和荒漠，南方有热带雨林，西部至西南部則有天山山脈、帕米爾高原、青藏高原、
喀喇崑崙山脈和喜馬拉雅山脈，东临太平洋。领海由渤海（内海）以及黄海、东海、南海三大边海组成[18]，水域面积约470万平方千米，分布有大小岛屿7600个。
中国疆域東至黑龙江省佳木斯市抚远市的黑瞎子岛中部，西至新疆境内的帕米尔高原，北至黑龙江省大兴安岭地区的漠河县，南至南海曾母暗沙。
中国是目前世界上人口最多的國家，約有14億人（不包括香港、澳门特别行政区及未实际管辖的台湾省）[20]，同時也是一个多民族国家，共有受到官方承認的
民族56個，其中汉族人口佔91.51%[21]。以普通话和规范汉字为国家通用语言文字，少数民族地区可使用自己民族的語言文字。自1986年实行九年義務教育制
度，就读公立学校的学生由政府提供其学费。
中国目前为世界第二大经济体，GDP总量仅次于美国。1978年改革開放後，中國成为經濟成長最快的經濟體之一[22][23]。当前，该国对外贸易额世界第一，
是世界上最大的商品出口國及第二大的進口國，依國內生產總值按購買力平價位列世界第一、而國際匯率則排名世界第二[24]。2018年，中国國內生產總值依
購買力平價為255,161.28亿美元[25]，同比2017年增长16,914.92亿美元，增长速度为7.10% ，增量同样居全球第一，增速在总量排名前20的国家中位居
第一[12]。2000年时中国人均GDP仅有959美元，但到2019年，依据國際匯率計算，中国GDP为14.36万亿美元，人均GDP已超过1万美元，而部分东部沿海省
市的人均國內生產總值更已经超过2~2.5万美元，对于众多的14亿人口大国能够达到这个数值，表明中国的经济体量和消费市场巨大。改革开放以来，贫困问题
已得到改善，然而國民贫富差距大等社會问题仍需解决[26][27]。
科技方面，中国在航天航空、装备制造业、高速鐵路、新能源、核技术、超级计算机、量子网络、人工智能、5G通訊等領域有较强实力，研發經費則位居世界第
二，也是世界第二个超万亿美元投入研发的国家[28]。國防預算為世界第二高每年超过1700亿美元的军费投入，擁有世界規模最大作战力量的常備部隊及三位
一體的核打擊能力并拥有在亚太地区局部优势的作战能力和拥有一支蓝水海军的作战力量。[29][30]
1949年中国共产党领导中国人民解放军在内战中取得优势，实际控制了中國大陸，并于同年10月1日宣布建立中华人民共和国和中央人民政府，与遷至台灣地區
的中華民國政府形成兩岸分治的格局至今。该国成立初期遵循和平共处五项原则的外交政策，在1971年取得在聯合國的中國代表權同时继承了原中華民國的联合
国安理会常任理事国地位后陆续加入了部分联合国其他专门机构，并广泛参与重要國際組織例如国际奥委会、亚太经合组织、二十国集团、世界贸易组织，并成
为了上海合作组织、金砖国家、一带一路、亚洲基础设施投资银行、区域全面经济伙伴关系协定等国际合作组织项目的发起国和创始国。随着该国的国际影响力
增强，已被许多国家、组织、智库视为世界重要的潜在超级大國之一和世界经济重要支柱。`
