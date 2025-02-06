packagemain

 import(
 "fmt"
 "net/http"
 "strconv"
 "os"
 "math"
)

 typeMass struct{
 Densityfloat64
 }

 //---ADDYOURCODE--


  //---BETWEENTHOSELINES--

   funcHandler(massVolumeMassVolume)http.HandlerFunc{
   returnfunc(whttp.ResponseWriter,r*http.Request){
   ifdimension,err:=strconv.ParseFloat(r.URL.Query().Get("dimension"),64);err==nil{
   weight :=massVolume.density() *massVolume.volume(dimension)
   w.Write([]byte(fmt.Sprintf("%.2f",math.Round(weight*100)/100)))
   return
   }
   w.WriteHeader(http.StatusBadRequest)
   }
   }
  
   funcmain() {
   port,err:=strconv.Atoi(os.Args[1])
   iferr!=nil{
   panic(err)
   }
  
   aluminiumSphere:=Sphere{Mass{Density:2.710}}
   ironCube:=Cube{Mass{Density:7.874}}
  
   http.HandleFunc("/aluminium/sphere",Handler(aluminiumSphere))
   http.HandleFunc("/iron/cube",Handler(ironCube))
  
   iferr:=http.ListenAndServe(fmt.Sprintf(":%d",port),nil);err!=nil{
   panic(err)
   }
   }