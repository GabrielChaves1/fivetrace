import { ComponentProps } from "react";
import { useTheme } from "./theme-provider";

export default function Grid({ ...props }: ComponentProps<'svg'>) {
  const { theme } = useTheme()

  return (
    <svg width="1159" height="832" viewBox="0 0 1159 832" className="fill-primary" {...props} xmlns="http://www.w3.org/2000/svg">
      <path d="M0.394084 0L0.394043 832M145.22 0L145.22 832M290.046 0L290.046 832M434.871 0L434.871 832M579.697 0L579.697 832M724.523 0L724.523 832M869.349 0L869.348 832M1014.17 0L1014.17 832M1159 0L1159 832M0 808.737H1158.61M0 677.842L1158.61 677.842M0 546.948L1158.61 546.948M0 416.053L1158.61 416.053M0 285.158L1158.61 285.158M0 154.263L1158.61 154.263M0 23.3685L1158.61 23.3685" stroke="url(#paint0_radial_310_39)" strokeOpacity="0.2" strokeWidth="0.6" />
      <defs>
        <radialGradient id="paint0_radial_310_39" cx="0" cy="0" r="1" gradientUnits="userSpaceOnUse" gradientTransform="translate(579.5 416) rotate(90) scale(416 579.5)">
          <stop stopColor={theme === "dark" ? "white" : "black"} />
          <stop offset="1" stopColor={theme === "dark" ? "white" : "black"} stopOpacity="0" />
        </radialGradient>
      </defs>
    </svg>
  )
}