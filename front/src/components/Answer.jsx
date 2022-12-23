import React from 'react'

function Answer({Answer,IsCorrect,removeAnswer,changeAnswer,changeIsCorrect,questionId,answerId}) {
  const handleAnswerChange = (e)=>changeAnswer(questionId,answerId,e.target.value)
  const handleIsCorrectChange = (e)=>changeIsCorrect(questionId,answerId,e.target.checked)
  
  return (
  
        <div className='answer-container'>
        <label className='text'>
    !
              <input value={Answer} onChange={handleAnswerChange} className='input'></input>
        </label>
          <label className='is_correct'>
              ✅
              <input className='input' checked={IsCorrect} onChange={handleIsCorrectChange} type='checkbox'></input>
            </label>
            <div className='remove' onClick={removeAnswer}>-</div>
      </div>
  )
}

export default Answer