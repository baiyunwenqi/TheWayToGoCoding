package packageStudy
import "regexp"
var ValidUrlFormat= regexp.MustCompile(`^(https?://)(([0-9]{1,3}\.){3}[0-9]{1,3}|([0-9a-z_!~*'()-]+\.)*([0-9a-z][0-9a-z-]{0,61})?[0-9a-z]\.[a-z]{2,6})(:[0-9]{1,4})?((/?)|(/[0-9a-z_!~*'().;?:@&=+$,%#-]+)+/?)$`)
func CheckValidUrl(url string ) bool {
	return  ValidUrlFormat.MatchString(url)
}
