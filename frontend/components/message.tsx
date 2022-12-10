import Image from 'next/image';
import { useState } from 'react';

interface Parameters {
   type: string;
   place?: string;
   mytext: string;
}

/**
* Notification Message
*/
const Message = ({ type, place, mytext }: Parameters): any => {
   const [check, setCheck] = useState<boolean>(true);
   const handleClick = () => {
      setCheck(false);
   };
   return (
      <>
         {check &&
            <div className="notification" >
               <div className={!place ? "right_bottom" : place} style={{ width: 800 }}>
                  <div className="notification_box">
                     <div className="text">
                        {type === "success" &&
                           <div className="success_icon">
                              <Image src="/notification_success.svg" alt="success" width={16} height={16} />
                           </div>
                        }
                        {type === "error" &&
                           <div className="error_icon">
                              <Image src="/notification_error.svg" alt="error" width={16} height={16} />
                           </div>
                        }
                        {mytext}
                     </div>
                     <div className="close_icon">
                        <a onClick={handleClick}>
                           <div>
                              <Image src="/notification_x.svg" alt="close" width={16} height={16} />
                           </div>
                        </a>
                     </div>
                  </div>
               </div>
            </div>
         }
      </>

   );

};
export default Message;