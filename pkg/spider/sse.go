package spider

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	pb "github.com/mineralres/protos/src/go/goshare"
	spiderpb "github.com/mineralres/protos/src/go/spider"
	"github.com/mineralres/goshare/pkg/util"
)

// SSE sse
type SSE struct {
}

func getURLContent(url, referer string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	if referer != "" {
		req.Header.Set("Referer", referer)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	return string(body), err
}

// OptionList 获取上证所网站的 50ETF个股期权列表
func (sse *SSE) OptionList() ([]*spiderpb.SSEStockOption, error) {
	const url = "http://query.sse.com.cn/commonQuery.do?jsonCallBack=jsonpCallback77327&isPagination=true&expireDate=&securityId=&sqlId=SSE_ZQPZ_YSP_GGQQZSXT_XXPL_DRHY_SEARCH_L&pageHelp.pageSize=10000&pageHelp.pageNo=1&pageHelp.beginPage=1&pageHelp.cacheSize=1&pageHelp.endPage=5&_=1531102881526"
	str, err := getURLContent(url, "http://www.sse.com.cn/assortment/options/disclo/preinfo/")
	if err != nil {
		return nil, err
	}
	var rsp struct {
		ActionErrors []int  `json:"actionErrors"`
		Locale       string `json:"locale"`
		IsPagination string `json:"isPagination"`
		PageHelp     struct {
			BeginPage int `json:"beginPage"`
			CacheSize int `json:"cacheSize"`
			Data      []struct {
				EXERCISE_PRICE        string `json:"EXERCISE_PRICE"` // 行权价
				UPDATE_VERSION        string `json:"UPDATE_VERSION"` //
				OPTION_TYPE           string `json:"OPTION_TYPE"`
				DAILY_PRICE_UPLIMIT   string `json:"DAILY_PRICE_UPLIMIT"` // 涨停价
				TIMESAVE              string `json:"TIMESAVE"`
				DELISTFLAG            string `json:"DELISTFLAG"`
				START_DATE            string `json:"START_DATE"`
				EXPIRE_DATE           string `json:"EXPIRE_DATE"`
				CONTRACT_UNIT         string `json:"CONTRACT_UNIT"`
				CALL_OR_PUT           string `json"CALL_OR_PUT"`
				LMTORD_MAXFLOOR       string `json:"LMTORD_MAXFLOOR"`
				DELIVERY_DATE         string `json:"DELIVERY_DATE"`
				CHANGEFLAG            string `json:"CHANGEFLAG"`
				MKTORD_MAXFLOOR       string `json:"MKTORD_MAXFLOOR"`
				UNDERLYING_TYPE       string `json:"UNDERLYING_TYPE"`
				DAILY_PRICE_DOWNLIMIT string `json:"DAILY_PRICE_DOWNLIMIT"`
				ROUND_LOT             string `json:"ROUND_LOT"`
				SECURITY_CLOSEPX      string `json:"SECURITY_CLOSEPX"`
				SETTL_PRICE           string `json:"SETTL_PRICE"`
				CONTRACT_SYMBOL       string `json:"CONTRACT_SYMBOL"`
				NUM                   string `json:"NUM"`
				CONTRACT_ID           string `json:"CONTRACT_ID"`
				MARGIN_RATIO_PARAM1   string `json:"MARGIN_RATIO_PARAM1"`
				MARGIN_RATIO_PARAM2   string `json:"MARGIN_RATIO_PARAM2"`
				LMTORD_MINFLOOR       string `json:"LMTORD_MINFLOOR"`
				MKTORD_MINFLOOR       string `json:"MKTORD_MINFLOOR"`
				END_DATE              string `json:"END_DATE"`
				PRICE_LIMIT_TYPE      string `json:"PRICE_LIMIT_TYPE"`
				EXERCISE_DATE         string `json:"EXERCISE_DATE"`
				MARGIN_UNIT           string `json:"MARGIN_UNIT"`
				SECURITY_ID           string `json:"SECURITY_ID"`
				SECURITYNAMEBYID      string `json:"SECURITYNAMEBYID"`
				CONTRACTFLAG          string `json:"CONTRACTFLAG"`
				UNDERLYING_CLOSEPX    string `json:"UNDERLYING_CLOSEPX"`
			} `json:"data"`
		} `json:"pageHelp"`
	}
	start := strings.Index(str, "(")
	str = str[start+1 : len(str)-1]
	err = json.Unmarshal([]byte(str), &rsp)
	if err != nil {
		return nil, err
	}
	var ret []*spiderpb.SSEStockOption
	for i := range rsp.PageHelp.Data {
		d := &rsp.PageHelp.Data[i]
		// log.Printf("合约编码[%s] 合约交易代码[%s] 合约简称[%s] 标的券名称及代码[%s] 类型[%s, %s] 行权价[%s] 合约单位[%s] 期权行权日[%s] 行权交收日[%s] 到期日[%s] 新挂[%s] 涨停价[%s] 跌停价[%s] 前结算价[%s] 调整[%s]",
		// 	d.SECURITY_ID, d.CONTRACT_ID, d.CONTRACT_SYMBOL, d.SECURITYNAMEBYID, d.CALL_OR_PUT, d.OPTION_TYPE, d.EXERCISE_PRICE, d.CONTRACT_UNIT, d.EXERCISE_DATE,
		// 	d.DELIVERY_DATE, d.EXPIRE_DATE, d.CHANGEFLAG, d.DAILY_PRICE_UPLIMIT, d.DAILY_PRICE_DOWNLIMIT, d.SETTL_PRICE, d.CHANGEFLAG)
		var op spiderpb.SSEStockOption
		op.ExercisePrice = d.EXERCISE_PRICE
		op.UpdateVersion = d.UPDATE_VERSION
		op.OptionType = d.OPTION_TYPE
		op.DailyPriceUpLimit = d.DAILY_PRICE_UPLIMIT
		op.TimeSave = d.TIMESAVE
		op.DELIST_Flag = d.DELISTFLAG
		op.StartDate = d.START_DATE
		op.ExpireDate = d.EXPIRE_DATE
		op.ContractUnit = d.CONTRACT_UNIT
		op.CallOrPut = d.CALL_OR_PUT
		op.LmtOrdMaxFloor = d.LMTORD_MAXFLOOR
		op.DeliveryDate = d.DELIVERY_DATE
		op.ChangeFlag = d.CHANGEFLAG
		op.MktOrdMaxFloor = d.MKTORD_MAXFLOOR
		op.UnderlyingClosePX = d.UNDERLYING_CLOSEPX
		op.UnderlyingType = d.UNDERLYING_TYPE
		op.DailyPriceDownLimit = d.DAILY_PRICE_DOWNLIMIT
		op.RoundLot = d.ROUND_LOT
		op.SecurityClosePX = d.SECURITY_CLOSEPX
		op.SettlPrice = d.SETTL_PRICE
		op.ContractSymbol = d.CONTRACT_SYMBOL
		op.Num = d.NUM
		op.ContractID = d.CONTRACT_ID
		op.MarginRatioParam1 = d.MARGIN_RATIO_PARAM1
		op.MarginRatioParam2 = d.MARGIN_RATIO_PARAM2
		op.LmtOrdMinFloor = d.LMTORD_MINFLOOR
		op.MktOrdMinFloor = d.MKTORD_MINFLOOR
		op.EndDate = d.END_DATE
		op.PriceLimitType = d.PRICE_LIMIT_TYPE
		op.ExerciseDate = d.EXERCISE_DATE
		op.MarginUnit = d.MARGIN_UNIT
		op.SecurityID = d.SECURITY_ID
		op.SecurityNameByID = d.SECURITYNAMEBYID
		op.ContractFlag = d.CONTRACTFLAG
		op.UnderlyingClosePX = d.UNDERLYING_CLOSEPX
		ret = append(ret, &op)
	}
	return ret, nil
}

// OptionInstrumentList 上证所ETF期权合约列表
func (sse *SSE) OptionInstrumentList() ([]*pb.Instrument, error) {
	list, err := sse.OptionList()
	if err != nil {
		return nil, err
	}
	// 期权行情
	var symbols []string
	var ret []*pb.Instrument
	for _, op := range list {
		ti := new(pb.Instrument)
		ti.Exchange = "SSE"
		ti.Symbol = op.SecurityID
		symbols = append(symbols, ti.Symbol)

		ti.Name = op.ContractSymbol
		ti.StrikePrice = util.ParseFloat(op.ExercisePrice)
		ti.UpperLimit = util.ParseFloat(op.DailyPriceUpLimit)
		ti.LowerLimit = util.ParseFloat(op.DailyPriceDownLimit)
		ti.PreSettlement = util.ParseFloat(op.SettlPrice)

		ud, _ := strconv.Atoi(time.Now().Format("20060102"))
		ti.TradingDay = int32(ud)
		ti.IsTrading = false
		ti.MaxMarketOrderVolume = int32(util.ParseInt(op.MktOrdMaxFloor))
		ti.MaxLimitOrderVolume = int32(util.ParseInt(op.LmtOrdMaxFloor))
		ti.MinMarketOrderVolume = 1
		ti.MinLimitOrderVolume = 1

		ti.ExpireDate = int32(util.ParseInt(op.ExpireDate))
		ti.StartDeliverDate = int32(util.ParseInt(op.DeliveryDate))
		ti.EndDeliverDate = ti.StartDeliverDate
		if ti.ExpireDate >= ti.TradingDay {
			ti.IsTrading = true
		}

		ti.PriceTick = 0.0001
		ti.ProductType = int32(pb.ProductType_SSE_ETF_OPTION)
		ti.Multiple = int32(util.ParseInt(op.ContractUnit))
		ti.Product = "SHOP"
		ti.DistinguishPositionTimeType = false

		if op.CallOrPut == "认购" {
			ti.CallOrPut = "call"
		} else if op.CallOrPut == "认沽" {
			ti.CallOrPut = "put"
		} else {
			panic("Invalid call put type")
		}
		ti.ExerciseDateType = "EUR"

		ret = append(ret, ti)
	}

	var sina Sina
	mdsList, err := sina.BatchGetSSEStockOptionTick(symbols)
	if err != nil {
		return ret, err
	}
	for i := range ret {
		ti := ret[i]
		for j := range mdsList {
			m := &mdsList[j]
			if m.Symbol == ti.Symbol && m.Exchange == ti.Exchange {
				ti.PrePosition = int32(m.Position)
				ti.PreClose = m.PreClose
				ti.PreSettlement = m.PreSettlement
			}
		}
	}
	return ret, nil
}

// ETF50OptionTQuote 上证所网站的T型报价,50ETF的
func (sse *SSE) ETF50OptionTQuote(month string) ([]*pb.OptionTQuoteItem, error) {
	var ret []*pb.OptionTQuoteItem
	url := fmt.Sprintf(`http://yunhq.sse.com.cn:32041/v1/sho/list/tstyle/510050_%s?select=contractid,code,last,chg_rate,presetpx,exepx,name,prev_close`, month)
	str, err := getURLContent(url, "http://www.sse.com.cn/assortment/options/price/")
	if err != nil {
		return ret, err
	}
	str = util.Decode(str)
	var xdata struct {
		Date  int             `json:"date"`
		Time  int             `json:"time"`
		Total int             `json:"total"`
		Begin int             `json:"begin"`
		End   int             `json:"end"`
		List  [][]interface{} `json:"list"`
	}
	err = json.Unmarshal([]byte(str), &xdata)
	if err != nil {
		log.Println("GetSSE50ETFOptionTQuote 获取T型报价错误", err, month)
		return ret, err
	}
	for i := range xdata.List {
		items := xdata.List[i]
		if len(items) != 8 {
			continue
		}
		contractid := items[0].(string)
		exercisePrice := items[5].(float64)
		name := items[6].(string)
		preClose := items[7].(float64)
		var st pb.SimpleTickForTQuote
		st.Exchange = "SSE"
		st.Symbol = items[1].(string)
		st.Price = items[2].(float64)
		st.UpDownRatio = items[3].(float64)
		if preClose > 0 {
			st.UpDownRatio = (st.Price - preClose) / preClose
			st.UpDownRatio *= 100
		}

		st.PreSettlementPrice = items[4].(float64)
		st.Name = name
		found := false
		exercisePriceFlag := strings.Trim(contractid, " ")

		if len(exercisePriceFlag) > 7 {
			exercisePriceFlag = exercisePriceFlag[7:]
		}
		for j := range ret {
			if ret[j].ExercisePriceFlag == exercisePriceFlag {
				found = true
				if strings.Contains(contractid, "C") {
					ret[j].Call = &st
				} else {
					ret[j].Put = &st
				}
			}
		}
		if !found {
			newItem := &pb.OptionTQuoteItem{}
			newItem.ExercisePrice = exercisePrice
			newItem.ExercisePriceFlag = exercisePriceFlag
			if strings.Contains(contractid, "C") {
				newItem.Call = &st
			} else {
				newItem.Put = &st
			}
			ret = append(ret, newItem)
		}
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ExercisePrice < ret[j].ExercisePrice
	})
	// log.Println(len(ret), month)

	return ret, nil
}

func downloadFile(url, target, referer string) error {
	client := &http.Client{
		Timeout: time.Second * 5,
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				deadline := time.Now().Add(30 * time.Second)
				c, err := net.DialTimeout(netw, addr, time.Second*5)
				if err != nil {
					return nil, err
				}
				c.SetDeadline(deadline)
				return c, nil
			},
		},
	}

	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	if referer != "" {
		req.Header.Set("Referer", referer)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	defer res.Body.Close()
	f, err := os.Create(target)
	if err != nil {
		log.Println("create file error ", err)
		return err
	}
	io.Copy(f, res.Body)
	f.Close()
	return nil
}

// StockList instrument list
func (sse *SSE) StockList(toFillPriceInfo bool) ([]*pb.Instrument, error) {
	excelFileName := fmt.Sprintf("SSE_STOCK_LIST_%s.xlsx", time.Now().Format("20060102"))
	if _, err := os.Stat(excelFileName); os.IsNotExist(err) {
		const url = "http://query.sse.com.cn/security/stock/downloadStockListFile.do?csrcCode=&stockCode=&areaName=&stockType=1"
		log.Println("下载文件", excelFileName, err)
		err := downloadFile(url, excelFileName, "http://www.sse.com.cn/market/stockdata/overview/monthly/")
		if err != nil {
			panic(err)
		}
		log.Println("下载完成")
	}
	file, err := os.Open(excelFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comment = '#' //可以设置读入文件中的注释符
	reader.Comma = '	'   //默认是逗号，也可以自己设置
	var list []*pb.Instrument
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("Error:", err)
			return list, err
		}
		if len(record) <= 5 {
			log.Println("格式有变", record, err, len(record))
			continue
		}

		var item pb.Instrument
		item.Exchange = "SSE"
		item.Symbol = record[2]
		item.Symbol = strings.TrimSpace(item.Symbol)
		opendate := strings.Replace(record[4], "-", "", -1)
		opendate = strings.Trim(opendate, " ")
		d, err := strconv.Atoi(opendate)
		if err != nil {
			// log.Println(err, record)
			continue
		}
		item.OpenDate = int32(d)

		item.Product = "SHA"
		item.PriceTick = 0.01
		item.Multiple = 1
		item.Name = util.Decode(record[3])
		item.UpdateTime = time.Now().Unix()
		item.ProductType = int32(pb.ProductType_STOCK)
		item.IsCloseTodayAllowed = false
		item.MaxLimitOrderVolume = 100000000
		item.MaxMarketOrderVolume = 100000000
		item.MinBuyVolume = 100
		item.MinLimitOrderVolume = 100
		item.MinMarketOrderVolume = 100
		item.MinSellVolume = 100
		item.IsTrading = true

		if strings.Contains(item.Name, "ST") || strings.Contains(item.Name, "*ST") {
			item.ProductClass = "ST"
		}

		ud, _ := strconv.Atoi(time.Now().Format("20060102"))
		item.TradingDay = int32(ud)

		list = append(list, &item)
	}
	if toFillPriceInfo {
		fillPriceInfo("SSE", list)
	}
	return list, nil
}
