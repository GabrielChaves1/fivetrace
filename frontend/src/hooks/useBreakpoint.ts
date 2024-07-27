import { useEffect, useState } from 'react';

const useBreakpoint = () => {
  const [isDesktop, setIsDesktop] = useState(window.innerWidth >= 1280);

  const handleResize = () => {
    setIsDesktop(window.innerWidth >= 1280);
  };

  useEffect(() => {
    window.addEventListener('resize', handleResize);
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  return isDesktop;
};

export default useBreakpoint;