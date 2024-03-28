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
	"fmt"
)

// ReviewSubmission defines model for ReviewSubmission.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmission
type ReviewSubmission struct {
	Attributes    *ReviewSubmissionAttributes    `json:"attributes,omitempty"`
	ID            string                         `json:"id"`
	Links         ResourceLinks                  `json:"links"`
	Relationships *ReviewSubmissionRelationships `json:"relationships,omitempty"`
	Type          string                         `json:"type"`
}

// ReviewSubmissionAttributes defines model for ReviewSubmissionAttributes.Attributes
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmission/attributes
type ReviewSubmissionAttributes struct {
	Platform      *string   `json:"platform,omitempty"`
	State         *string   `json:"state,omitempty"`
	SubmittedDate *DateTime `json:"submittedDate,omitempty"`
}

// ReviewSubmissionRelationships defines model for ReviewSubmission.Relationships
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmission/relationships
type ReviewSubmissionRelationships struct {
	App                      *Relationship      `json:"app,omitempty"`
	AppStoreVersionForReview *Relationship      `json:"appStoreVersionForReview,omitempty"`
	Items                    *PagedRelationship `json:"items,omitempty"`
	LastUpdatedByActor       *Relationship      `json:"lastUpdatedByActor,omitempty"`
	SubmittedByActor         *Relationship      `json:"submittedByActor,omitempty"`
}

// ReviewSubmissionResponse defines model for ReviewSubmissionResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionresponse
type ReviewSubmissionResponse struct {
	Data     ReviewSubmission                   `json:"data"`
	Included []ReviewSubmissionResponseIncluded `json:"included,omitempty"`
	Links    DocumentLinks                      `json:"links"`
}

// ReviewSubmissionsResponse defines model for ReviewSubmissionsResponse.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionsresponse
type ReviewSubmissionsResponse struct {
	Data     []ReviewSubmission                 `json:"data"`
	Included []ReviewSubmissionResponseIncluded `json:"included,omitempty"`
	Links    PagedDocumentLinks                 `json:"links"`
	Meta     *PagingInformation                 `json:"meta,omitempty"`
}

// ReviewSubmissionResponseIncluded is a heterogenous wrapper for the possible types that can be returned
// in a ReviewSubmissionResponse.
type ReviewSubmissionResponseIncluded included

// reviewSubmissionUpdateRequest defines model for reviewSubmissionUpdateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionupdaterequest/data
type reviewSubmissionUpdateRequest struct {
	Attributes *reviewSubmissionUpdateRequestAttributes `json:"attributes,omitempty"`
	ID         string                                   `json:"id"`
	Type       string                                   `json:"type"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionupdaterequest/data/attributes
type reviewSubmissionUpdateRequestAttributes struct {
	Canceled  *bool `json:"canceled,omitempty"`
	Submitted *bool `json:"submitted,omitempty"`
}

// ListSubmissionsForAppQuery are query options for ListSubmissionsForAppQuery
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_review_submissions_for_an_app
type ListSubmissionsForAppQuery struct {
	FieldsReviewSubmissionItems []string `url:"fields[reviewSubmissionItems],omitempty"`
	FieldsReviewSubmissions     []string `url:"fields[reviewSubmissions],omitempty"`
	FilterApp                   []string `url:"filter[app],omitempty"`
	FilterState                 []string `url:"filter[state],omitempty"`
	Include                     []string `url:"include,omitempty"`
	Limit                       int      `url:"limit,omitempty"`
	LimitItems                  int      `url:"limit[items],omitempty"`
}

// ReadReviewSubmissionQuery are query options for ReadReviewSubmissionQuery
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_review_submission_information
type ReadReviewSubmissionQuery struct {
	FieldsReviewSubmissionItems []string `url:"fields[reviewSubmissionItems],omitempty"`
	FieldsReviewSubmissions     []string `url:"fields[reviewSubmissions],omitempty"`
	Include                     []string `url:"include,omitempty"`
	LimitItems                  int      `url:"limit[items],omitempty"`
}

// reviewSubmissionCreateRequest defines model for reviewSubmissionCreateRequest.
//
// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissioncreaterequest/data
type reviewSubmissionCreateRequest struct {
	Attributes    reviewSubmissionCreateRequestAttributes    `json:"attributes,omitempty"`
	Relationships reviewSubmissionCreateRequestRelationships `json:"relationships"`
	Type          string                                     `json:"type"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/reviewsubmissionupdaterequest/data/attributes
type reviewSubmissionCreateRequestAttributes struct {
	Platform Platform `json:"platform,omitempty"`
}

// appStoreVersionSubmissionCreateRequestRelationships are attributes for AppStoreVersionSubmissionCreateRequest
//
// https://developer.apple.com/documentation/appstoreconnectapi/appstoreversionsubmissioncreaterequest/data/relationships
type reviewSubmissionCreateRequestRelationships struct {
	App relationshipDeclaration `json:"app"`
}

// ListReviewSubmissionsForApp finds and lists review submissions for all apps in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/list_review_submissions_for_an_app
func (s *ReviewSubmissionsService) ListReviewSubmissionsForApp(ctx context.Context, params *ListSubmissionsForAppQuery) (*ReviewSubmissionsResponse, *Response, error) {
	res := new(ReviewSubmissionsResponse)

	resp, err := s.client.get(ctx, "reviewSubmissions", params, res)

	return res, resp, err
}

// GetReviewSubmission find review submission in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/read_review_submission_information
func (s *ReviewSubmissionsService) GetReviewSubmission(ctx context.Context, id string, params *ReadReviewSubmissionQuery) (*ReviewSubmissionResponse, *Response, error) {
	url := fmt.Sprintf("reviewSubmissions/%s", id)
	res := new(ReviewSubmissionResponse)
	resp, err := s.client.get(ctx, url, params, res)

	return res, resp, err
}

// CreateReviewSubmission create review submission for app and platform in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/create_a_review_submission
func (s *ReviewSubmissionsService) CreateReviewSubmission(ctx context.Context, platform Platform, appID string) (*ReviewSubmissionResponse, *Response, error) {
	req := reviewSubmissionCreateRequest{
		Attributes: reviewSubmissionCreateRequestAttributes{
			Platform: platform,
		},
		Relationships: reviewSubmissionCreateRequestRelationships{
			App: relationshipDeclaration{
				Data: RelationshipData{
					ID:   appID,
					Type: "apps",
				},
			},
		},
		Type: "reviewSubmissions",
	}

	res := new(ReviewSubmissionResponse)
	resp, err := s.client.post(ctx, "reviewSubmissions", newRequestBody(req), res)

	return res, resp, err
}

// UpdateReviewSubmission update review submission for app and platform in App Store Connect.
//
// https://developer.apple.com/documentation/appstoreconnectapi/modify_a_review_submission
func (s *ReviewSubmissionsService) UpdateReviewSubmission(ctx context.Context, id string, canceled *bool, submitted *bool) (*ReviewSubmissionResponse, *Response, error) {
	req := reviewSubmissionUpdateRequest{
		ID:   id,
		Type: "reviewSubmissions",
	}

	if canceled != nil || submitted != nil {
		req.Attributes = &reviewSubmissionUpdateRequestAttributes{}
		if canceled != nil {
			req.Attributes.Canceled = canceled
		}
		if submitted != nil {
			req.Attributes.Submitted = submitted
		}
	}

	url := fmt.Sprintf("reviewSubmissions/%s", id)

	res := new(ReviewSubmissionResponse)
	resp, err := s.client.patch(ctx, url, newRequestBody(req), res)

	return res, resp, err
}

// App returns the App stored within, if one is present.
func (i *ReviewSubmissionResponseIncluded) App() *App {
	return extractIncludedApp(i.inner)
}

// AppStoreVersion returns the AppStoreVersion stored within, if one is present.
func (i *ReviewSubmissionResponseIncluded) AppStoreVersion() *AppStoreVersion {
	return extractIncludedAppStoreVersion(i.inner)
}
