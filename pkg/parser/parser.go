package parser

import "time"

func ParseTimeString(timeStr string) (time.Time, error) {
    layout := "2006-01-02T15:04:05Z" 
    parsedTime, err := time.Parse(layout, timeStr)
    if err != nil {
        return time.Time{}, err
    }
    return parsedTime, nil
}
