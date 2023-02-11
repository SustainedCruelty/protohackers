import Control.Concurrent
import Control.Monad (unless)
import Data.ByteString
import Network.Socket
import Network.Socket.ByteString (recv, sendAll)

main :: IO ()
main = do
  sock <- socket AF_INET Stream 0
  setSocketOption sock ReuseAddr 1
  bind sock (SockAddrInet 8082 0)
  listen sock 5
  mainLoop sock
  close sock

mainLoop :: Socket -> IO ()
mainLoop sock = do
  (conn, addr) <- accept sock
  forkIO (sendEcho conn)
  mainLoop sock

sendEcho :: Socket -> IO ()
sendEcho conn = do
  msg <- recv conn 1024
  unless (Data.ByteString.null msg) $ do
    sendAll conn msg
    sendEcho conn
    close conn