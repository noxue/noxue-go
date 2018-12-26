/**
 * @author 刘荣飞 yes@noxue.com
 * @date 2018/12/26 23:55
 */

 package utils

type Error string

func (e Error) Error() string {
	return string(e)
}

func CheckErr(err error) {
	if err != nil {
		panic(Error(err.Error()))
	}
}
