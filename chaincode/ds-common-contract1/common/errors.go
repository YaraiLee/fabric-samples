/**
 * @Author: liyalei
 * @Description:
 * @Version:
 * @Date: 2020/6/23 10:50 上午
 */
package common

import "errors"

var (
	ErrChainCodeArgsNum   = errors.New("chain code argv num err")
	ErrChainCodeFn        = errors.New("chain code fn err")
	ErrChainQueryNotExist = errors.New("chain query content not exist")
	ErrOwnerSetFail       = errors.New("not admin identity")
	ErrIllegalIdentity    = errors.New("illegal identity not the owner")
	ErrNotFound           = errors.New("data does not exist")
	ErrParamInvalid       = errors.New("chain code param invalid")
)
