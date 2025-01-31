package data

import (
	"fmt"
	"time"
)

func (f *FeedData) add(item FeedItem) {
	f.items[item.ID] = item
}

func (f *FeedData) GetIDs() []string {
	var retVal []string

	for key := range f.items {
		retVal = append(retVal, key)
	}

	return retVal
}

func (f *FeedData) GetItem(id string) (FeedItem, error) {
	if item, ok := f.items[id]; ok {
		return item, nil
	}
	return FeedItem{}, fmt.Errorf("could not find item with ID: %s", id)
}

func (f *FeedData) GetAllItems() []FeedItem {
	var retVal []FeedItem
	now := time.Now()

	for _, item := range f.items {
		retVal = append(retVal, item)
	}

	f.LastDisplayTime = now

	return retVal
}

func (f *FeedData) GetNewItems() []FeedItem {
	var retVal []FeedItem
	now := time.Now()

	for _, item := range f.items {
		if item.CreatedAt.After(f.LastDisplayTime) {
			retVal = append(retVal, item)
		}
	}

	f.LastDisplayTime = now

	return retVal
}
