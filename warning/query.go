package warning

import (
	"github.com/TerrexTech/go-eventstore-models/model"
	"github.com/TerrexTech/go-mongoutils/mongo"
)

func Query(collection *mongo.Collection, event *model.Event) *model.KafkaResponse {
	switch event.ServiceAction {
	// case "timestamp":
	// 	return queryTimestamp(collection, event)
	case "count":
		return queryCount(collection, event)
	default:
		return queryWarning(collection, event)
	}
}

// // Query handles "query" events.
// func Query(collection *mongo.Collection, event *model.Event) *model.KafkaResponse {
// 	filter := map[string]interface{}{}

// 	err := json.Unmarshal(event.Data, &filter)
// 	if err != nil {
// 		err = errors.Wrap(err, "Query: Error while unmarshalling Event-data")
// 		log.Println(err)
// 		return &model.KafkaResponse{
// 			AggregateID:   event.AggregateID,
// 			CorrelationID: event.CorrelationID,
// 			Error:         err.Error(),
// 			ErrorCode:     InternalError,
// 			UUID:          event.UUID,
// 		}
// 	}

// 	if len(filter) == 0 {
// 		err = errors.New("blank filter provided")
// 		err = errors.Wrap(err, "Query")
// 		log.Println(err)
// 		return &model.KafkaResponse{
// 			AggregateID:   event.AggregateID,
// 			CorrelationID: event.CorrelationID,
// 			Error:         err.Error(),
// 			ErrorCode:     InternalError,
// 			UUID:          event.UUID,
// 		}
// 	}

// 	result, err := collection.Find(filter)
// 	if err != nil {
// 		err = errors.Wrap(err, "Query: Error in DeleteMany")
// 		log.Println(err)
// 		return &model.KafkaResponse{
// 			AggregateID:   event.AggregateID,
// 			CorrelationID: event.CorrelationID,
// 			Error:         err.Error(),
// 			ErrorCode:     DatabaseError,
// 			UUID:          event.UUID,
// 		}
// 	}

// 	resultMarshal, err := json.Marshal(result)
// 	if err != nil {
// 		err = errors.Wrap(err, "Query: Error marshalling Warning Delete-result")
// 		log.Println(err)
// 		return &model.KafkaResponse{
// 			AggregateID:   event.AggregateID,
// 			CorrelationID: event.CorrelationID,
// 			Error:         err.Error(),
// 			ErrorCode:     InternalError,
// 			UUID:          event.UUID,
// 		}
// 	}

// 	return &model.KafkaResponse{
// 		AggregateID:   event.AggregateID,
// 		CorrelationID: event.CorrelationID,
// 		Result:        resultMarshal,
// 		UUID:          event.UUID,
// 	}
// }
