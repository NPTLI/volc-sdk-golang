// Code generated by protoc-gen-volcengine-sdk
// source: SnapshotManageService
// DO NOT EDIT!

package live

import (
	"encoding/json"
	"net/http"
	"net/url"

	"google.golang.org/protobuf/encoding/protojson"

	"github.com/pkg/errors"

	"github.com/volcengine/volc-sdk-golang/service/live/models/request"
	"github.com/volcengine/volc-sdk-golang/service/live/models/response"
)

// DescribeCDNSnapshotHistory
/*
 * @param *request.DescribeCDNSnapshotHistoryRequest
 * @return *response.DescribeCDNSnapshotHistoryResponse, int, error
 */
func (p *LIVE) DescribeCDNSnapshotHistory(req *request.DescribeCDNSnapshotHistoryRequest) (*response.DescribeCDNSnapshotHistoryResponse, int, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	respBody, status, err := p.Json("DescribeCDNSnapshotHistory", url.Values{}, string(body))
	output := &response.DescribeCDNSnapshotHistoryResponse{}
	unmarshaler := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
	errUnmarshal := unmarshaler.Unmarshal(respBody, output)
	if err != nil || status != http.StatusOK {
		// if exist http err,check whether the respBody's type is defined struct,
		// if it is ,
		// return struct,
		// otherwise return nil body
		// if httpCode is not 200,check whether the respBody's type is defined struct,
		// if it is ,
		// use errorCode as err and return struct,
		// otherwise use respBody string as error and return
		if errUnmarshal != nil || len(output.GetResponseMetadata().GetError().GetCode()) == 0 {
			if err == nil {
				err = errors.New(string(respBody))
			}
			return nil, status, err
		} else {
			return output, status, errors.New(output.GetResponseMetadata().GetError().GetCode())
		}
	}
	return output, status, nil
}
