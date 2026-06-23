param(
    [int]$Display = 2,
    [ValidateSet(0,90,180,270)]
    [int]$Rotate = 0
)

Add-Type @"
using System;
using System.Runtime.InteropServices;

public class DisplayUtil {
  [StructLayout(LayoutKind.Sequential, CharSet=CharSet.Ansi)]
  public struct DEVMODE {
    [MarshalAs(UnmanagedType.ByValTStr, SizeConst=32)]
    public string dmDeviceName;
    public short dmSpecVersion;
    public short dmDriverVersion;
    public short dmSize;
    public short dmDriverExtra;
    public int dmFields;
    public int dmPositionX;
    public int dmPositionY;
    public int dmDisplayOrientation;
    public int dmDisplayFixedOutput;
    public short dmColor;
    public short dmDuplex;
    public short dmYResolution;
    public short dmTTOption;
    public short dmCollate;
    [MarshalAs(UnmanagedType.ByValTStr, SizeConst=32)]
    public string dmFormName;
    public short dmLogPixels;
    public int dmBitsPerPel;
    public int dmPelsWidth;
    public int dmPelsHeight;
    public int dmDisplayFlags;
    public int dmDisplayFrequency;
  }

  [DllImport("user32.dll")]
  public static extern int EnumDisplaySettings(string deviceName, int modeNum, ref DEVMODE devMode);

  [DllImport("user32.dll")]
  public static extern int ChangeDisplaySettingsEx(string deviceName, ref DEVMODE devMode, IntPtr hwnd, int flags, IntPtr lParam);
}
"@

$device = "\\.\DISPLAY$Display"

$dm = New-Object DisplayUtil+DEVMODE
$dm.dmSize = [Runtime.InteropServices.Marshal]::SizeOf($dm)

if ([DisplayUtil]::EnumDisplaySettings($device, -1, [ref]$dm) -eq 0)
{
    throw "找不到屏幕：$device"
}

$old = $dm.dmDisplayOrientation
$new = @{ 0 = 0; 90 = 1; 180 = 2; 270 = 3 }[$Rotate]

if (($old % 2) -ne ($new % 2))
{
    $tmp = $dm.dmPelsWidth
    $dm.dmPelsWidth = $dm.dmPelsHeight
    $dm.dmPelsHeight = $tmp
}

$dm.dmDisplayOrientation = $new
$dm.dmFields = 0x180080 # orientation + width + height

$result = [DisplayUtil]::ChangeDisplaySettingsEx($device, [ref]$dm, [IntPtr]::Zero, 1, [IntPtr]::Zero)

if ($result -ne 0)
{
    throw "旋转失败，错误码：$result"
}

"已旋转 $device 到 $Rotate 度"
