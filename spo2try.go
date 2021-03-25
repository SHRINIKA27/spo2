package main

import (
    "fmt"
    "encoding/csv"
    "time"
        "os"
    "strconv"
    "math/big"
    "math/rand"



)

var (

    ID string   
    SPO2 string
)

var remainder = big.NewInt(0);
var pk = big.NewInt(0);
var SPO2e = big.NewInt(0);
var IDe = big.NewInt(0);

const MAX int = 10000000;
var get_key int =0;
var arr_key int =0;
var arr_key1 int =3;
var get_keyvalues int =0;
var a1 int =0;    
var b1 int =0;
var j int =0;
var result string;

var spo2_range  = make([]*big.Int, 0)
var pkey  = make([]*big.Int, 0)

var key_hard = make([]int,0)
var keylist = make([]int,0)
var list = make([]int,0)
 func Encryption_Initialization() {
           
            //prime get  
            get_key=get_key+1;
            get_keyvalues=get_key+1;
            key_generation := new(big.Int)  
            
            key_generation.Mul(big.NewInt(int64(list[get_key])),big.NewInt(int64(list[get_keyvalues])));
            pkey=append(pkey,big.NewInt(int64(list[get_keyvalues])));
            //random key
            var result int =key_hard[arr_key];
            arr_key+=1;
            
            //power 
            power_values := new(big.Int)  

            power_values.Exp((big.NewInt(int64(list[get_keyvalues]))),big.NewInt(int64(list[get_key])),nil);
            //final power
            final_power:= new(big.Int)  

            final_power.Mul(power_values,big.NewInt(int64(result)));

            //mod
            remainder.Mod(final_power,key_generation);
    }

    func Encryption_process(val string) (ncrypted_unit *big.Int) {

            id1:= new(big.Int)  
            ncrypted_unit = new(big.Int)  
            id1,ok:= id1.SetString(val, 10)
            if !ok {
             fmt.Println("SetString: error")
             return
            }
            ncrypted_unit.Add(id1,remainder);
            return            
    }

func Encryption() {
    
    //random number generation

    get_key =1500+rand.Intn(400);
        var i int 
        for  i = 0;i<MAX;i++ {
                      
                    var key int = 150+rand.Intn(200000);
                    if key>0 {
                        
                        key_hard = append(key_hard,key);
                        }       
        }
    //  prime number gneration

    var beg=1500
    var end=20000 
    var n int 
    var j int 
            for  n = beg; n <= end; n++ {
            var prime = true;
            for  j = 2; j <= n / 2; j++ {
                if n % j == 0 && n != j {
                    prime = false;
                }
            }
            if prime {
                list = append(list,n);
            }
            }    


    Encryption_Initialization();
    
    IDe=Encryption_process(ID) 
    
    fmt.Println("---------------------Encrypted values----------------------")
    fmt.Print("ID :")
    fmt.Println(IDe)
    fmt.Print("SPO2 :")
    SPO2e = Encryption_process(SPO2) 
    
    fmt.Println(SPO2e)
    fmt.Println("------------------------------------------------------------")

    spo2_r := "91";

    spo2_re:= new(big.Int) 
    spo2_re=Encryption_process(spo2_r)



    spo2_range = append(spo2_range,spo2_re);

    //insertion into gorti table
}
func Analysis()  {


       var c1 =0;
       result="Person safe"


       c1=spo2_range[b1].Cmp(SPO2e)
       if (c1==0 || c1==1) {

        result="Person should be admitted"
        
       }
     



       

}
func Decryption_process(id1 *big.Int) (n1 *big.Int) {
        
    n1 = new(big.Int) 
    n1.Mod(id1,pkey[b1]);
    return 
}

func Decryption() {
    
        
        IDd:= new(big.Int) 
        IDd =Decryption_process(IDe)
        
        
        SPO2d:= new(big.Int) 
        SPO2d =Decryption_process(SPO2e)
        


    fmt.Println()
    fmt.Println("---------------Decrypted values---------------")
    fmt.Print("ID :")
    fmt.Println(IDd)
    fmt.Print("SPO2 :")
    fmt.Println(SPO2d)
    fmt.Print("RESULT :")
    fmt.Println(result)
    fmt.Println("----------------------------------------------")
    fmt.Println()
    b1++;
}


func main() {

    rand.Seed(time.Now().UnixNano())
    csvFile, err := os.Open("oxitable.csv")
    if err != nil {
        fmt.Println(err)
    }

        fmt.Println("Successfully Opened CSV file")
    defer csvFile.Close()
    
    csvLines, err := csv.NewReader(csvFile).ReadAll()
    if err != nil {
        fmt.Println(err)
    }   
    var i int
    i = 1
    for _, line := range csvLines {

            ID = strconv.Itoa(i)
            SPO2 = line[0]

        

        Encryption()
        Analysis()
        Decryption()
        i++;
    
    }
    
    
}
