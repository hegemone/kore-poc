module Kore
  module Comm
    class PlatformAdapter < Kore::Machinery::Base
      TRIGGER_PREFIX = '!'

      def initialize(engine)
        super()
        @engine = engine
      end

      def message_received(originator, raw)
        self.log.debug "msg: [#{raw}]"

        if self.is_cmd(raw)
          self.log.debug "Command matched, parsing..."
          begin
            m = self.parse(raw)
            m.originator = originator
          rescue Exception => e
            self.log.warn "Error occurred while parsing message"
            self.log.warn raw
            self.log.warn e.message
          else
            @engine.route_ingress(m)
          end
        end
      end

      #############################################################
      # Interface methods to be overriden
      #############################################################
      def send_message(emsg)
        raise "platform adapter #{self.name} must implement \"send_message\" method"
      end

      def name
        raise 'platform adapters must implement "name" method'
      end

      # Must be non-blocking!
      def listen
        raise "platform adapter #{self.name} must implement \"listen\" method"
      end
      #############################################################

      def is_cmd(raw)
        !!(Regexp.new("^#{TRIGGER_PREFIX}\\S*($| )").match(raw))
      end

      def parse(raw)
        # TODO: TEST ALL THE THINGS

        # Message structure:
        # <TRIGGER_PREFIX><PLUGIN> <CMD?> <CONTENT?>
        # Example message types:
        # "!bacon" -- no cmd
        # "!satellite ping" -- no content
        # "!droplet start awesomesauce" -- full
        if raw[0] != TRIGGER_PREFIX
          raise "attempting to parse command without trigger prefix: #{TRIGGER_PREFIX}"
        end

        snipped = raw[1..raw.length]
        tmp = snipped.split(" ")
        plugin = tmp.shift
        cmd = if tmp.length != 0 then tmp.shift else "" end
        content = tmp.join(" ")

        Kore::Comm::IngressMessage.new({
          plugin: plugin,
          cmd: cmd,
          content: content,
          raw: raw
        })
      end
    end
  end
end
