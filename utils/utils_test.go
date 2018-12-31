/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/30 20:22
 */
package utils

import (
	"fmt"
	"testing"
)

func TestGenerateVerifyCode(t *testing.T) {
	fmt.Println(GenerateVerifyCode())
}
