import React from 'react'
import {Link} from "react-router-dom";
import { useNavigate } from 'react-router-dom';
import {UserContext} from '../../App'

function Login(){
  
  return (
<article className="rounded-xl border-2 border-gray-100 bg-white">
        <div className="flex items-center gap-4 p-4 sm:p-6 lg:p-8">
          

            <img
              alt="Speaker"
              src="https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8NHx8YXZhdGFyfGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=800&q=60"
              className="h-14 w-14 rounded-full object-cover"
            />


          <div>
            <h3 className="font-medium sm:text-lg">
                Nodev918
            </h3>

            <p className="line-clamp-2 text-sm text-gray-700">
              Lorem ipsum dolor, sit amet consectetur adipisicing elit. Accusamus,
              accusantium temporibus iure delectus ut totam natus nesciunt ex?
              Ducimus, enim.
            </p>
          </div>
          <div className="flex justify-end">
            <strong className="-mb-[2px] -me-[2px] inline-flex items-center gap-1 rounded-ee-xl rounded-ss-xl bg-green-600 px-3 py-1.5 text-white">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-4 w-4"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                strokeWidth="2"
              >
                <path
                  strokeLinecap="round"
                  strokeLinejoin="round"
                  d="M9 12l2 2 4-4M7.835 4.697a3.42 3.42 0 001.946-.806 3.42 3.42 0 014.438 0 3.42 3.42 0 001.946.806 3.42 3.42 0 013.138 3.138 3.42 3.42 0 00.806 1.946 3.42 3.42 0 010 4.438 3.42 3.42 0 00-.806 1.946 3.42 3.42 0 01-3.138 3.138 3.42 3.42 0 00-1.946.806 3.42 3.42 0 01-4.438 0 3.42 3.42 0 00-1.946-.806 3.42 3.42 0 01-3.138-3.138 3.42 3.42 0 00-.806-1.946 3.42 3.42 0 010-4.438 3.42 3.42 0 00.806-1.946 3.42 3.42 0 013.138-3.138z"
                />
              </svg>

              <span className="text-[10px] font-medium sm:text-xs">Solved!</span>
            </strong>
          </div>
        </div>


      </article>
  )
}

export default Login