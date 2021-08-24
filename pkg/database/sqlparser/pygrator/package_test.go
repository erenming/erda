// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pygrator_test

import (
	"path/filepath"
	"testing"

	"github.com/erda-project/erda/pkg/database/sqlparser/pygrator"
)

const (
	pythonFilename = "20210701-my-feature.py"
)

const pythonText = `from datetime import datetime
import django.db.models


class Questions(django.db.models.Model):
    """
    generated by erda-cli migrate init
    """

    question_text = django.db.models.CharField(max_length=20)
    pub_date = django.db.models.DateTimeField("date published")

    class Meta:
        db_table = "polls_questions"


class Choice(django.db.models.Model):
    id = django.db.models.BigIntegerField()
    choice_text = django.db.models.CharField(max_length=20)

    class Meta:
        db_table = "polls_choice"


class Second(django.db.models.Model):
    choice_text = django.db.models.CharField()


def add_question():
    questions = Questions()
    questions.question_text = "the first question"
    questions.pub_date = datetime.now()
    questions.save(force_insert=True)


def add_choice():
    all_choice = Choice.objects.filter()
    print("all choice:", all_choice)
    for choice in all_choice:
        print(choice.id, choice.choice_text)
    choice = Choice()
    choice.choice_text = "this is a choice this is a choice this is a choice this is a choice this is a choice"
    choice.save()


entries: [callable] = [
    add_question,
    add_choice,
]
`

type textFile struct {
	Name string
	Data []byte
}

func (f textFile) GetName() string {
	return filepath.Base(f.Name)
}

func (f textFile) GetData() []byte {
	return f.Data
}

func TestPackage_Make(t *testing.T) {
	var p pygrator.Package
	p.DeveloperScript = textFile{
		Name: pythonFilename,
		Data: []byte(pythonText),
	}
	p.Requirements = []byte(pygrator.BaseRequirements)
	p.Settings = settings
	p.Commit = true

	if err := p.Make(); err != nil {
		t.Fatalf("failed to p.Make: %v", err)
	}

}

func TestPackage_Remove(t *testing.T) {
}

func TestPackage_Run(t *testing.T) {

}
