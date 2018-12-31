/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/31 10:17
 */
package srv

import (
	"fmt"
	"testing"
)

func TestCaptcha_Create(t *testing.T) {

	fmt.Println(SrvCaptcha.Create())
}
