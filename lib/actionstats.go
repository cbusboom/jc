//------------------------------------------------------------------------------
// This library class provides action statistics.
//
// Action times are recorded in the statistics via the AddAction method.
// Action averages can be retrieved via the GetStats method.
//
// Each of these methods are described below:
//------------------------------------------------------------------------------
package actionstats

import (
    "encoding/json"
    "sync"
)

// This tracks the statistics data
type actionData struct {
    Action string `json:"action"`
    Avg int       `json:"avg"`
    time int
    count int
}

// This is the ActionStats class
type ActionStats struct {
    mutex sync.RWMutex
    stats map[string]*actionData
}

//------------------------------------------------------------------------------
// This is the constructor for a new ActionStats object
//------------------------------------------------------------------------------
func NewActionStats() ActionStats {
    stats := ActionStats{}
    stats.stats = make(map[string]*actionData)
    return stats
}

//------------------------------------------------------------------------------
// This function accepts a JSON serialized string of the form below 
// and maintains an average time for each action.
// Below are three sample inputs:
// 
// 1) {"action":"jump","time":100}
// 2) {"action":"run","time":75}
// 3) {"action":"jump","time":200}
//------------------------------------------------------------------------------
func (a ActionStats) AddAction(js string) error {
    type actionJSON struct {
        Action string `json:"action"`
        Time int      `json:"time"`
    }
    var data actionJSON
    err := json.Unmarshal([]byte(js),&data)
    if err != nil {
        return err
    }
    a.mutex.Lock()
    if _, ok := a.stats[data.Action]; ok {
        a.stats[data.Action].time += data.Time
        a.stats[data.Action].count++
        a.stats[data.Action].Avg = a.stats[data.Action].time / a.stats[data.Action].count
    } else {
        a.stats[data.Action] = &actionData{data.Action,data.Time,data.Time,1}
    }
    a.mutex.Unlock()
    return nil
}

//------------------------------------------------------------------------------
// This function returns a JSON array of the average time for each action 
// that has been recorded by the addAction function. The output after the 
// 3 sample inputs (see AddAction above) would be:
//
//   [{"action":"jump","avg":150},{"action":"run","avg":75}]
//------------------------------------------------------------------------------
func (a ActionStats) GetStats() string {
    data := make([]actionData, 0, len(a.stats))

    a.mutex.RLock()
    for _, stat := range a.stats {
        data =  append(data, *stat)
    }
    a.mutex.RUnlock()

    jsonStats, _ := json.Marshal(data)

    return string(jsonStats)
}
