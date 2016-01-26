/*
   Copyright (C) 2016 AHaymond

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program. If not, see {http://www.gnu.org/licenses/}.
*/

package jira

// customfield_10602 (effort)
type Effort struct {
	Self  string `json:"self,omitempty"`
	Value string `json:"value,omitempty"`
	Id    string `json:"id,omitempty"`
}

// customfield_10604 (team)
type Team struct {
	Self  string `json:"self,omitempty"`
	Value string `json:"value,omitempty"`
	Id    string `json:"id,omitempty"`
}
