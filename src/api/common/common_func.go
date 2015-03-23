package common
import (
    //"encoding/json"
    "fmt"
    "os"
    "github.com/codeskyblue/go-sh"
)

func DisplayJson(obj_json map[string]interface{}){
    fmt.Println("----------------------parse start------------------------")
    for k, v := range obj_json  {
        switch v2 := v.(type) {
        case string:
            fmt.Println(k, "is string", v2)
        case int:
            fmt.Println(k, "is int ", v2)
        case bool:
            fmt.Println(k, "is bool", v2)
        case []interface{}:
            fmt.Println(k, "is array", v2)
            for i, iv := range v2 {
                    fmt.Println(i,iv)
            }
        case map[string]interface{}:
            fmt.Println(k, "is map")
            DisplayJson(v2)
        default:
            fmt.Println(k, "is another type not handle yet")
        }
    }
    fmt.Println("----------------------parse end------------------------")
}

func TransferFileSSH(strSrcFile string,strDestFile string)(ret int,err string){
    session := sh.NewSession()
    session.ShowCMD = true    
    err1:= session.Call("scp",strSrcFile,strDestFile)
    if err1 != nil {
        fmt.Println("transfer remote file faild error:", err1)
        return 1,"transfer remote file faild error"
    }

    return 0,"ok"
}

func CreateRemotePath(strServerIP string,strPath string)(ret int,err string){
    session := sh.NewSession()
    session.ShowCMD = true    
    err1:= session.Call("ssh",strServerIP, "mkdir ",strPath)
    if err1 != nil {
        fmt.Println("exec remote shell faild error:", err1)
        return 1,"exec remote shell faild error"
    }

    return 0,"ok"
}
func SaveFile(strFileName string,strData string)(ok bool){
    f, err := os.Create(strFileName)
    if err != nil {
        fmt.Println("create file faild error:", err)
        return false
    }
    _, err_w := f.Write([]byte(strData))
    if err_w != nil {
        fmt.Println("Server start faild error:", err_w)
        return false
    }
    return true
}
