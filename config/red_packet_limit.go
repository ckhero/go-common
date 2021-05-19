/**
 *@Description
 *@ClassName red_packet_limit
 *@Date 2021/5/18 上午10:58
 *@Author ckhero
 */

package config

type RedPacketLimit struct {
	Amount uint64 `json:"amount"`
	RecvNum uint64 `json:"recvNum"`
}


func GetRedPacketLimit() *RedPacketLimit {
	return appConfig.RedPacketLimit
}