/**
Copyright (C) 2024 vfadc.

This file is part of asc-go, a package for working with Apple's
App Store Connect API.

asc-go is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

asc-go is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with asc-go.  If not, see <http://www.gnu.org/licenses/>.
*/

package asc

import (
	"context"
)

// ReviewSubmissionItem defines model for ReviewSubmissionItem.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionitem
type ReviewSubmissionItem struct {
	Attributes    *ReviewSubmissionItemAttributes    `json:"attributes,omitempty"`
	ID            string                             `json:"id"`
	Links         ResourceLinks                      `json:"links"`
	Relationships *ReviewSubmissionItemRelationships `json:"relationships,omitempty"`
	Type          string                             `json:"type"`
}

// ReviewSubmissionItemAttributes defines model for ReviewSubmissionItemAttributes.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionitem/attributes
type ReviewSubmissionItemAttributes struct {
	State *string `json:"state,omitempty"`
}

// ReviewSubmissionItemRelationships defines model for ReviewSubmissionItem.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionitem/relationships
type ReviewSubmissionItemRelationships struct {
	AppCustomProductPageVersion *Relationship `json:"appCustomProductPageVersion,omitempty"`
	AppEvent                    *Relationship `json:"appEvent,omitempty"`
	AppStoreVersion             *Relationship `json:"appStoreVersion,omitempty"`
	AppStoreVersionExperiment   *Relationship `json:"appStoreVersionExperiment,omitempty"`
	AppStoreVersionExperimentV2 *Relationship `json:"appStoreVersionExperimentV2,omitempty"`
}

// ReviewSubmissionItemResponse defines model for ReviewSubmissionItemResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionitemsresponse
type ReviewSubmissionItemResponse struct {
	Data     ReviewSubmissionItem                   `json:"data"`
	Included []ReviewSubmissionItemResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                          `json:"links"`
}

// ReviewSubmissionItemsResponse defines model for ReviewSubmissionItemsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionitemsresponse
type ReviewSubmissionItemsResponse struct {
	Data     []ReviewSubmissionItem                 `json:"data"`
	Included []ReviewSubmissionItemResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                     `json:"links"`
	Meta     *PagingInformation                     `json:"meta,omitempty"`
}

// ReviewSubmissionItemResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a ReviewSubmissionItemResponse.
type ReviewSubmissionItemResponseIncluded included

// reviewSubmissionItemUpdateRequest defines model for reviewSubmissionItemUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionitemupdaterequest
type reviewSubmissionItemUpdateRequest struct {
	Attributes *reviewSubmissionItemUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                       `json:"id"`
	Type       string                                       `json:"type"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionitemupdaterequest/data/attributes
type reviewSubmissionItemUpdateRequestAttributes struct {
	Removed  *bool `json:"removed,omitempty"`
	Resolved *bool `json:"resolved,omitempty"`
}

// reviewSubmissionItemCreateRequest defines model for reviewSubmissionItemCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionitemcreaterequest/data
type reviewSubmissionItemCreateRequest struct {
	Relationships reviewSubmissionItemCreateRequestRelationships `json:"relationships"`
	Type          string                                         `json:"type"`
}

// appStoreVersionSubmissionCreateRequestRelationships are attributes for AppStoreVersionSubmissionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmissioncreaterequest/data/relationships
type reviewSubmissionItemCreateRequestRelationships struct {
	AppCustomProductPageVersion *relationshipDeclaration `json:"appCustomProductPageVersion,omitempty"`
	AppEvent                    *relationshipDeclaration `json:"appEvent,omitempty"`
	AppStoreVersion             *relationshipDeclaration `json:"appStoreVersion,omitempty"`
	AppStoreVersionExperiment   *relationshipDeclaration `json:"appStoreVersionExperiment,omitempty"`
	AppStoreVersionExperimentV2 *relationshipDeclaration `json:"appStoreVersionExperimentV2,omitempty"`
	ReviewSubmission            relationshipDeclaration  `json:"reviewSubmission"`
}

// CreateReviewSubmissionItem finds and lists review submissions for all apps in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/post_v1_reviewsubmissionitems
func (s *ReviewSubmissionsService) CreateReviewSubmissionItem(ctx context.Context, appStoreVersionID *string, reviewSubmissionID string) (*ReviewSubmissionItemResponse, *Response, error) {
	req := reviewSubmissionItemCreateRequest{
		Relationships: reviewSubmissionItemCreateRequestRelationships{
			ReviewSubmission: relationshipDeclaration{
				Data: RelationshipData{
					ID:   reviewSubmissionID,
					Type: "reviewSubmissions",
				},
			},
		},
		Type: "reviewSubmissionItems",
	}
	if appStoreVersionID != nil {
		req.Relationships.AppStoreVersion = &relationshipDeclaration{
			Data: RelationshipData{
				ID:   *appStoreVersionID,
				Type: "appStoreVersions",
			},
		}
	}

	res := new(ReviewSubmissionItemResponse)
	resp, err := s.client.post(ctx, "reviewSubmissionItems", newRequestBody(req), res)

	return res, resp, err
}
