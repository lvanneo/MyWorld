package main
 
import(
    "encoding/json"
    "io/ioutil"
    "log"
    "net/http"
)
 
type WeatherInfoJson struct{
    Weatherinfo WeatherinfoObject
}
 
type WeatherinfoObject struct{
    City string
    CityId string
    Temp string
    WD string
    WS string
    SD string
    WSE string
    Time string
    IsRadar string
    Radar string
}
 
func main(){
    log.SetFlags(log.LstdFlags|log.Lshortfile)
    resp,err:=http.Get("http://www.weather.com.cn/data/sk/101010100.html")
    if err!=nil{
        log.Fatal(err)
    }
 
    defer resp.Body.Close()
    input,err:=ioutil.ReadAll(resp.Body)
 
    var jsonWeather WeatherInfoJson
    json.Unmarshal(input,&jsonWeather)
    log.Printf("Results:%v\n",jsonWeather)
 
    log.Println(jsonWeather.Weatherinfo.City)
    log.Println(jsonWeather.Weatherinfo.WD)
    log.Println(jsonWeather.Weatherinfo.WS)
    log.Println(jsonWeather.Weatherinfo.SD)
    log.Println(jsonWeather.Weatherinfo.WSE)
    log.Println(jsonWeather.Weatherinfo.Time)
    
 
    //ioutil.WriteFile("wsk101010100.html",input,0644)
}