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

// SubmissionService handles communication with submission-related methods of the App Store Connect API
//
// https://developer.apple.com/documentation/appstoreconnectapi/advertising_identifier_idfa_declarations
// https://developer.apple.com/documentation/appstoreconnectapi/app_store_review_details
// https://developer.apple.com/documentation/appstoreconnectapi/app_store_review_attachments
// https://developer.apple.com/documentation/appstoreconnectapi/app_store_version_submissions
type ReviewSubmissionsService service
