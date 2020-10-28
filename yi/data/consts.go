package data

import "github.com/nomos/go-lokas/protocol"

type PERIOD protocol.Enum

const (
	PERIOD_MIN1 PERIOD = iota + 1
	PERIOD_MIN3
	PERIOD_MIN5
	PERIOD_MIN15
	PERIOD_MIN30
	PERIOD_MIN60
	PERIOD_H1
	PERIOD_H2
	PERIOD_H3
	PERIOD_H4
	PERIOD_H6
	PERIOD_H8
	PERIOD_H12
	PERIOD_DAY1
	PERIOD_DAY3
	PERIOD_WEEK1
	PERIOD_MONTH1
	PERIOD_YEAR1
)

func (this PERIOD) String() string {
	switch this {
	case PERIOD_MIN1:
		return "1m"
	case PERIOD_MIN3:
		return "3m"
	case PERIOD_MIN5:
		return "5m"
	case PERIOD_MIN15:
		return "15m"
	case PERIOD_MIN30:
		return "30m"
	case PERIOD_MIN60:
		return "60m"
	case PERIOD_H1:
		return "1h"
	case PERIOD_H2:
		return "2h"
	case PERIOD_H3:
		return "3h"
	case PERIOD_H4:
		return "4h"
	case PERIOD_H6:
		return "6h"
	case PERIOD_H8:
		return "8h"
	case PERIOD_H12:
		return "12h"
	case PERIOD_DAY1:
		return "1d"
	case PERIOD_DAY3:
		return "3d"
	case PERIOD_WEEK1:
		return "1w"
	case PERIOD_MONTH1:
		return "1M"
	case PERIOD_YEAR1:
		return "1y"
	default:
		return "none"
	}
}
