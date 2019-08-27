package cronicleapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 事件状态
type EventStatusObject struct {
	EventObject
	CommonRsp
	Hostname   string  `json:"hostname"`
	Complete   Nbool   `json:"complete"`
	Source     string  `json:"source"`
	Progress   float64 `json:"progress"`
	TimeStart  int64   `json:"time_start"`
	TimeEnd    int64   `json:"time_end"`
	// second
	Elapsed    int64   `json:"elapsed"`
	// cpi mem log pid ....
}

// 事件对象
type EventObject struct {
	// event ID
	Id          string    `json:"id,omitempty"`
	//Timing   
	Timing      *TimingObject `json:"timing"`   
	Title       string    `json:"title"`
	Enabled     Nbool      `json:"enabled"`
	Algo        string    `json:"algo,omitempty"`
	ApiKey      string    `json:"api_key,omitempty"`
	CatchUp     Nbool     `json:"catch_up,omitempty"`
	Category    string    `json:"category"`

	// timestamp
	Created     int64     `json:"created,omitempty"`
	Modified    int64     `json:"modified,omitempty"`
	// plugin define
	Plugin      string    `json:"plugin"`
	Params      json.RawMessage `json:"params,omitempty"`
	
	// second
	Timeout     int32     `json:"timeout,omitempty"`
	Retries     int32     `json:"retries,omitempty"`
	// second
	RetryDelay  int32     `json:"retry_delay,omitempty"`


	Target      string    `json:"target"`

	Timezone    string    `json:"timezone,omitempty"`
	
	// success event ID
	Chain       string    `json:"chain,omitempty"`
	// error event ID
	ChainError  string    `json:"chain_error,omitempty"`
	CpuLimit    int32     `json:"cpu_limit,omitempty"`
	CpuSustain  int32     `json:"cpu_sustain,omitempty"`
	Detached    Nbool      `json:"detached,omitempty"`
	MaxChildren int32     `json:"max_children,omitempty"`
	MemoryLimit int32     `json:"memory_limit,omitempty"`
	MemorySustain int32   `json:"memory_sustain,omitempty"`
	Multiplex   Nbool      `json:"multiplex,omitempty"`
	Notes       string    `json:"notes,omitempty"`
	NotifyFail  string    `json:"notify_fail,omitempty"`
	NotifySuccess string  `json:"notify_success,omitempty"`
	Queue       Nbool      `json:"queue,omitempty"`
	QueueMax    int32     `json:"queue_max,omitempty"`
	Stagger     int32     `json:"stagger,omitempty"`
	Username    string    `json:"username,omitempty"`
	WebHook     string    `json:"web_hook,omitempty"`
}

type Nbool bool

func (b *Nbool) UnmarshalJSON(data []byte) error {
    asString := string(data)
    if asString == "1" || asString == "true" {
        *b = true
    } else if asString == "0" || asString == "false" {
        *b = false
    } else {
        return errors.New(fmt.Sprintf("Boolean unmarshal error: invalid input %s", asString))
    }
    return nil
}

func (b Nbool) MarshalJSON() ([]byte, error) {
	s := "0"
	if b {
		s = "1"
	}
    return []byte(s), nil
}