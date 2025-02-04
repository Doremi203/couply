import React from 'react'
import styles from './homePage.module.css'
import Like from './like'


export default function HomePage() {

      
  return (
    <div>
        <div> Couply</div>

        <div className={styles.personCard}>
            <img src='../public/photo1.png' width='350px' height='530px' />
            </div>        
        
        <div className={styles.likeCircle}> 
            <Like />
        </div>


    </div>
        
  )
}
