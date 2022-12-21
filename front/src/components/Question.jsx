import React from 'react'
import Answer from './Answer'

function Question({Question,Timer,Answers,id,removeQuestion,addAnswer,removeAnswer,changeQuestion , changeTimer, ...rest}) {
    const answersComponents =  Answers.map((answer,ansId)=><Answer removeAnswer={()=>removeAnswer (id,ansId)} questionId={id} answerId={ansId} Answer={answer.Answer} key={'ans_'+id+'_'+ansId} IsCorrect={answer.IsCorrect} {...rest} ></Answer>)
    const handleChangeQuestion = (e)=> changeQuestion(id,e.target.value)
    const handleChangeTimer = (e)=> changeTimer(id,e.target.value)
    return (
    <div className='question-container'>
      <div className='remove' onClick={() => removeQuestion(id)}>⌫</div>
          <label className='question-label'>  
              ⍰
              <textarea value={Question} onChange={handleChangeQuestion} className='question'></textarea>
          </label>
          <label className='question-label'>
            ⏲
              <select className='timer' onChange={handleChangeTimer}>
                <option value={30}>30"</option>
                <option value={60}>1'</option>
                <option value={120}>2'</option>
            </select>
          </label>
            {answersComponents}
        <div className='add' onClick={()=>addAnswer(id)}>+</div>
    </div>
  )
}

export default Question