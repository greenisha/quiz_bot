import React, { useState } from 'react'
import Question from './Question'

function Quiz() {
    const defaultAnswer = {Answer:"",IsCorrect:false}
    const defaultQuestion = {
        Question:"Почём рыба?",
        Timer:30,
      Answers: [{...defaultAnswer}, { ...defaultAnswer }]
    }
    const [Questions, setQuestions] = useState([{...defaultQuestion}])
    const addQuestion = ()=>{setQuestions([...Questions,{...defaultQuestion}])}
    const removeQuestion = (id) =>{setQuestions(Questions.filter((_,idx)=>idx!=id))}
    const removeAnswer = (questionId,answerId) =>{
      let question = Questions[questionId];
      question.Answers = question.Answers.filter((_, idx) => idx != answerId)
      let questions=[...Questions]
      questions[questionId]=question
      setQuestions(questions);
    }
    const changeQuestion = (questionId,Question)=>{
      let question = Questions[questionId];
      question.Question = Question;
      let localquestions = [...Questions]
      localquestions[questionId] = question
      setQuestions(localquestions);
    }
    const addAnswer = (questionId) => {
      let question = Questions[questionId];
      question.Answers = [...question.Answers,defaultAnswer]
      let questions = [...Questions]
      questions[questionId] = question
      setQuestions(questions);

    }
    const changeAnswer = (questionId,answerId,answer) =>{
      const question = Questions[questionId];
      question.Answers[answerId].Answer=answer;
      let questions = [...Questions]
      questions[questionId] = question
      setQuestions(questions);
    }
    const changeIsCorrect = (questionId,answerId,isCorrect) =>{
      const question = Questions[questionId];
      question.Answers[answerId].IsCorrect=isCorrect;
      let questions = [...Questions]
      questions[questionId] = question
      setQuestions(questions);
    }
    const changeTimer = (questionId,timer)=>{
      let question = Questions[questionId];
      question.Timer = timer;
      let questions = [...Questions]
      questions[questionId] = question
      setQuestions(questions);
      
    }
    const sendQuestion =(_) =>{
      console.log(Questions);
      fetch('/api/addQuiz',{
        method:'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({Questions:Questions}),
      })
      .then((response)=>response.json())
      .then((data)=>{
        console.log('Success:',data)
      })
      .catch((error)=>{
        console.error('Error:',error);
      });
    }
    const QuestionsComponents = Questions.map((question,id)=><Question changeTimer={changeTimer} changeAnswer={changeAnswer} changeIsCorrect={changeIsCorrect} changeQuestion={changeQuestion} removeAnswer={removeAnswer} addAnswer={addAnswer} removeQuestion={removeQuestion} id={id} key={'quest_'+id} Answers={question.Answers} Timer={question.Timer} Question={question.Question}></Question>)
 // console.log(QuestionsComponents);
    return (<>
    <div className=''>{QuestionsComponents}</div>
      <div className='add' onClick={addQuestion}>⊕</div>
      <div className='send' onClick={sendQuestion}>✉</div>
    </>
  )
}

export default Quiz