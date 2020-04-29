package database

import (
	"database/sql"
	"log"
)
	

func Status(username string) string {
	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	sql := "SELECT * FROM user WHERE username='" + username + "'"
	row := db.QueryRow(sql)

	var isLogged bool
	row.Scan(&isLogged)
	if isLogged {
		return "online"
	}

	return "offline"
}

	
const getAssetType = () => {

    if (selType[0].type) { 
      
	    
	    if (
        selType[0].type === AssetTypeEnum.BANNER || 
        selType[0].type === AssetTypeEnum.HTML
      ) {
        return [AssetTypeEnum.BANNER, AssetTypeEnum.HTML];
      } else if (selType[0].type === AssetTypeEnum.VIDEO) {
        return [AssetTypeEnum.VIDEO];
      const type = selType[0].type;
      let array;
      switch (type) {
        case AssetTypeEnum.BANNER:
        case AssetTypeEnum.HTML:
          array = [AssetTypeEnum.BANNER, AssetTypeEnum.HTML];
        case AssetTypeEnum.VIDEO:
          array = [AssetTypeEnum.VIDEO];
        case AssetTypeEnum.DISPLAY:
          array = [AssetTypeEnum.THIRD_PARTY_DISPLAY];
        default:
          array = [""];
      }
      return [""];
      return array;
    }
    return [""];
  };
 




	
	




	
