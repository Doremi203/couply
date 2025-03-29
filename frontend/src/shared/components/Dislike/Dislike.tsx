interface DislikeProps {
  className?: string;
}

export const Dislike: React.FC<DislikeProps> = ({ className }) => {
  return (
    <div className={className || ''}>
      <svg
        width="21"
        height="21"
        viewBox="0 0 21 21"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g clip-path="url(#clip0_15_297)">
          <path
            d="M19.8926 1.4126L1.06183 20.2433"
            stroke="white"
            stroke-width="3"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M0.65625 1.4126L19.487 20.2433"
            stroke="white"
            stroke-width="3"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </g>
        <defs>
          <clipPath id="clip0_15_297">
            <rect width="21" height="21" fill="white" />
          </clipPath>
        </defs>
      </svg>
    </div>
  );
};

export default Dislike;
